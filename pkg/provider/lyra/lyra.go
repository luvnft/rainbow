// Copyright 2022 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package lyra

import (
	"fmt"
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/teal-finance/emo"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

var log = emo.NewZone("lyr")

const (
	optimismrpc        = "https://opt-mainnet.g.alchemy.com/v2/6_IOOvszkG_h71cZH3ybdKrgPPwAUx6m" // "https://mainnet.optimism.io"
	name               = "Lyra"
	lyraRegistry       = "0xF5A0442D4753cA1Ea36427ec071aa5E786dA5916"
	optionMarketViewer = "0xEAf788AD8abd9C98bA05F6802a62B8DbC673D76B"
	QuoterAddress      = "0xea83ee73eB397c5974CB6b5220AE0A32fbE48B2B"
	oneOption          = 1
)

type Provider struct{}

func (Provider) Name() string {
	return name
}

func (Provider) Options() ([]rainbow.Option, error) {
	options := []rainbow.Option{}
	client, err := ethclient.Dial(optimismrpc)
	if err != nil {
		return nil, log.Error("Lyra ethclient", err).Err()
	}

	address := common.HexToAddress(lyraRegistry)
	registry, err := NewLyrar(address, client)
	if err != nil {
		return nil, log.Error("Lyra registry", err).Err()
	}

	// DO NOT use make() here!
	// we don't want to have 0x00...0 initialize here.
	// and we don't know how many market there will be.
	markets := []common.Address{}
	var i int64
	var market common.Address
	err = nil
	for ; err == nil; i++ {
		market, err = registry.OptionMarkets(&bind.CallOpts{}, big.NewInt(i))
		if err == nil {
			markets = append(markets, market)
		}
	}
	if len(markets) == 0 {
		log.Error("registry.OptionMarkets return empty array")
	}

	sum := 0
	viewer, err := NewLyrap(common.HexToAddress(optionMarketViewer), client)
	if err != nil {
		return nil, log.Error("optionMarketViewer", err).Err()
	}
	quoter, err := NewLyraq(common.HexToAddress(QuoterAddress), client)
	if err != nil {
		return nil, log.Error("quoter contract", err).Err()
	}

	for i := 0; i < len(markets); i++ {
		marketAddresses, err := viewer.MarketAddresses(&bind.CallOpts{}, markets[i])
		if err != nil {
			return nil, log.Error("MarketAddresses", err).Err()
		}
		baseAsset := Asset(marketAddresses.BaseAsset)

		boards, err := viewer.GetLiveBoards(&bind.CallOpts{}, markets[i])
		if err != nil {
			return nil, log.Error("GetLiveBoards", err).Err()
		}

		for _, b := range boards {
			sum += len(b.Strikes)

			for _, s := range b.Strikes {
				callPut, err := process(s, b, baseAsset, quoter)
				if err != nil {
					return nil, log.Error("process", err).Err()
				}
				options = append(options, callPut...)
			}
		}
	}

	log.Info("Lyra total markets", sum)

	return options, nil
}

func process(s OptionMarketViewerStrikeView, b OptionMarketViewerBoardView, asset string, quoter *Lyraq) ([]rainbow.Option, error) {
	options := []rainbow.Option{}

	call := rainbow.Option{
		Name:          "",
		Type:          "CALL",
		Asset:         asset,
		Expiry:        expiration(b.Expiry),
		ExchangeType:  "DEX",
		Chain:         "Ethereum",
		Layer:         "L2",
		LayerName:     "Optimism",
		Provider:      name,
		QuoteCurrency: "USD", // sUSD but anyway
		Bid:           nil,
		Ask:           nil,
		Strike:        rainbow.ToFloat(s.StrikePrice, rainbow.DefaultEthereumDecimals),
	}
	put := rainbow.Option{
		Name:          "",
		Type:          "PUT",
		Asset:         asset,
		Expiry:        expiration(b.Expiry),
		ExchangeType:  "DEX",
		Chain:         "Ethereum",
		Layer:         "L2",
		LayerName:     "Optimism",
		Provider:      name,
		QuoteCurrency: "USD", // sUSD but anyway
		Bid:           nil,
		Ask:           nil,
		Strike:        rainbow.ToFloat(s.StrikePrice, rainbow.DefaultEthereumDecimals),
	}

	call.Name = call.OptionName()
	put.Name = put.OptionName()

	call.URL = url(call, s.StrikeId)
	put.URL = url(put, s.StrikeId)

	// Market IV = board IV (baseIV) * Skew
	call.MarketIV = rainbow.ToFloat(b.BaseIv, rainbow.DefaultEthereumDecimals) *
		rainbow.ToFloat(s.Skew, rainbow.DefaultEthereumDecimals)
	// keep only 5 decimals (IV is already a % so it can be shown as XX.XXX%)
	call.MarketIV = math.Floor(call.MarketIV*100000) / 100000
	put.MarketIV = call.MarketIV

	bidasks, err := getBidsAsks(s.StrikeId, b.Market, oneOption, quoter)
	if err != nil {
		return options, log.Error("getBidsAsks", err).Err()
	}

	call.Bid = append(call.Bid, rainbow.Order{
		Price: bidasks[3],
		Size:  float64(oneOption),
	})

	call.Ask = append(call.Ask, rainbow.Order{
		Price: bidasks[0],
		Size:  float64(oneOption),
	})
	put.Bid = append(put.Bid, rainbow.Order{
		Price: bidasks[4],
		Size:  float64(oneOption),
	})

	put.Ask = append(put.Ask, rainbow.Order{
		Price: bidasks[1],
		Size:  float64(oneOption),
	})
	options = append(options, call, put)

	return options, err
}

func getBidsAsks(strikeId *big.Int, market common.Address, amount int, quoter *Lyraq) ([]float64, error) {
	// TODO check what iteration really mean in Lyra
	// I know they use 1 :-)
	premium, _, err := quoter.FullQuotes(&bind.CallOpts{}, market, strikeId, common.Big1,
		rainbow.IntToEthereumUint256(amount, 18))
	if err != nil {
		return nil, log.Error("FullQuotes market=", market, "strikeId=", strikeId, err).Err()
	}
	return rainbow.ToFLoat(premium, rainbow.DefaultEthereumDecimals), nil
}

func Asset(address common.Address) string {
	// those asset are part of Synthetix so if it's not recognized
	// it is an unknwow synthetic asset.

	switch {
	case address.String() == sETH:
		return "sETH"
	case address.String() == sBTC:
		return "sBTC"
	case address.String() == sSOL:
		return "sSOL"
	case address.String() == sLINK:
		return "sLINK"
	default:
		log.Warn("Lyra Unknown token:", address.String())
		return "LLLL"
	}
}

// TODO check function on their frontend.
func url(o rainbow.Option, strikeId *big.Int) string {
	base := "https://app.lyra.finance/trade"
	asset := strings.ToLower(o.Asset[1:])
	// TODO if they include asset with decimal, modify this
	// most likely example: SOL, LINK
	strike := fmt.Sprintf("%.f", rainbow.ToFloat(strikeId, 0))
	t := strings.ToLower(o.Type)

	return base + "/" + asset + "/" + strike + "/" + t
}

func expiration(e *big.Int) string {
	seconds := e.Int64()
	expiryTime := time.Unix(seconds, 0).UTC()
	return expiryTime.Format("2006-01-02 15:04:05")
}
