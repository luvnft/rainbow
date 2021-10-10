// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package psyoptions

import (
	"context"
	"fmt"
	"math/big"

	"github.com/streamingfast/solana-go"
	"github.com/streamingfast/solana-go/programs/serum"
	"github.com/streamingfast/solana-go/rpc"

	"github.com/teal-finance/rainbow"
)

const (
	listMarketsURL   = "wss://api.psyoptions.io/v1/graphql"
	PsyQuoteCurrency = "USDC"
	mainnet          = "https://solana-api.projectserum.com" //https://api.mainnet-beta.solana.com"
	devnet           = "https://api.devnet.solana.com"
	Expiration       = "2021-10-29 23:59:59"
)

func Instruments(coin string) []string {
	if coin == "ETH" {
		return []string{
			"8A493gU55NfS4fCjDoLAiN57zPzWf6QQw31QQf1fd6iX",
			"5pHcU2Gz8eCMwynLvz1AHSFoFbKkeTc7ufeqeG4spb99",
			"8qgAVVE6eeJ9u32LvJupgW6NyNWPPFFK4xnXfGtDNeP4",
			"FfkVR6ha6N9fcefGxfDFP97xL54Pc5ipiwnt1uuTpuyX",
			"TLAzx53rb6pEDT3oWitTzg6YNe7BTERLDzujMdU8RQx",
			"9A1m1JXgnMrdq3JkqEYx9rK3CcYm9EHDLeA6sn1hH3PP",
			"HYmPvo8szh62QVaAfUAXR1eppvCfouUPpH68yE87UYmy",
			"8fFcWuVaZSKoge4DCpMcNrR5nNXF2pbXCfBUxkMomgr5",
		}
	} else if coin == "BTC" {
		return []string{
			"8fhiAYm41RwtiX8WusCSpY617GWPt2LwUnCQcEeer78o",
			"6at26sVk8vTYtLh4YDKvje4PDdgFJsNHHyoGw87WNszP",
			"2gKrDsubuvYKxTkWdT5b44Qdd9QoBRTQQebUoQNnsesw",
			"7W2LGEDpitCoXLC5xhzjUKiE4NnNkgoAstM2EyFt7MaS",
			"9ugAWZCSgUKjL11fJE9Zjn4QVTdTkAkSLgPu9ZC8mcfD",
			"ACdjLA5wPk31eUEqra9BFQ3MTXbHqZfdM1TRQPX8Hi28",
			"2q5f1H8xT3tsBzQhwZC3BKnbKMb44fTuDGamZ6xUdZz2",
			"DvohGwDZR9Z2siWBj2Xhgxd1qRScVCpywL3EoRbpon3p",
		}
	}
	return []string{}
}

func InstrumentsFromAllMarkets() (r []rainbow.Options, err error) {
	//instruments := append(Instruments("ETH"), Instruments("BTC")...)
	instruments := GetInstruments()
	client := rpc.NewClient(mainnet)

	for _, i := range instruments {

		pubKey := solana.MustPublicKeyFromBase58(i.SerumMarketAddress)

		ctx := context.TODO()
		out, err := serum.FetchMarket(ctx, client, pubKey)
		if err != nil {
			panic(err)
		}

		bids, totalBids, err := BidsAsksToOffers(ctx, out, client, out.Market.GetBids(), false, "BUY")
		if err != nil {
			panic(err)
		}
		asks, totalAsks, err := BidsAsksToOffers(context.TODO(), out, client, out.Market.GetAsks(), true, "SELL")
		if err != nil {
			panic(err)
		}
		fmt.Println("total ", totalAsks+totalBids)
		offers := append(bids, asks...)

		o := rainbow.Options{
			Name:         i.Name(),
			Type:         i.Type(),
			Asset:        i.Asset(),
			Expiry:       Expiration,
			Strike:       i.Strike(),
			ExchangeType: "DEX",
			Chain:        "Solana",
			Layer:        "L1",
			Provider:     "PsyOptions",
			Offers:       offers,
		}
		r = append(r, o)

	}
	return r, err
}

// I don't really need the totalsize but I am keeping it since it was in the original func
// ASKs on the top so desc=True & side="SELL"
// BIDS down so desc=False @ side ="BUY"
func BidsAsksToOffers(ctx context.Context, market *serum.MarketMeta, cli *rpc.Client, address solana.PublicKey, desc bool, side string) (offers []rainbow.Offer, totalSize float64, err error) {
	var o serum.Orderbook
	if err := cli.GetAccountDataIn(ctx, address, &o); err != nil {
		return nil, 0, fmt.Errorf("getting orderbook: %w", err)
	}

	limit := 20
	levels := [][]*big.Int{}

	err = o.Items(desc, func(node *serum.SlabLeafNode) error {
		quantity := big.NewInt(int64(node.Quantity))
		price := node.GetPrice()
		if len(levels) > 0 && levels[len(levels)-1][0].Cmp(price) == 0 {
			current := levels[len(levels)-1][1]
			levels[len(levels)-1][1] = new(big.Int).Add(current, quantity)
		} else if len(levels) == limit {
			return fmt.Errorf("done")
		} else {
			levels = append(levels, []*big.Int{price, quantity})
		}
		return nil
	})
	if err != nil {
		//TODO fail more graciously
		panic(err)
	}

	for _, level := range levels {
		p := market.PriceLotsToNumber(level[0])
		price, _ := p.Float64()
		q := market.BaseSizeLotsToNumber(level[1])
		qty, _ := q.Float64()
		totalSize += qty

		offers = append(offers,
			rainbow.Offer{
				Price:         price,
				Quantity:      qty,
				QuoteCurrency: PsyQuoteCurrency,
				Side:          side,
			},
		)
	}

	return offers, totalSize, nil
}