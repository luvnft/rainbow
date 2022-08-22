// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package deribit

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/teal-finance/garcon"
	"github.com/teal-finance/rainbow/pkg/rainbow"
)

type Provider struct {
	ar AdaptiveRate
}

func (Provider) Name() string {
	return "Deribit"
}

// adaptiveMinSleepTime is to limit the request rate to the Deribit API.
const adaptiveMinSleepTime = 1 * time.Millisecond

// https://www.deribit.com/kb/deribit-rate-limits
// "Each sub-account has a rate limit of 100 in a burst or 20 requests per second".
func sleep(i int) {
	if i%50 == 0 {
		time.Sleep(1 * time.Second)
	}
}

// Hour at which the options expires = 8:00 UTC.
const Hour = 8

// maxBytesToRead prevents wasting memory/CPU when receiving an abnormally huge response from Deribit API.
const maxBytesToRead = 2_000_000

func (p *Provider) Options() ([]rainbow.Option, error) {
	if p.ar.Name == "" {
		p.ar = NewAdaptiveRate("Deribit", adaptiveMinSleepTime)
	}

	instruments, err := p.query("BTC")
	if err != nil {
		return nil, err
	}

	optionsBTC, err := p.fillOptions(instruments, 5)
	if err != nil {
		return nil, err
	}

	instruments, err = p.query("ETH")
	if err != nil {
		return nil, err
	}

	optionsETH, err := p.fillOptions(instruments, 5)
	if err != nil {
		return nil, err
	}

	instruments, err = p.query("SOL")
	if err != nil {
		return nil, err
	}

	optionsSOL, err := p.fillOptions(instruments, 5)
	if err != nil {
		return nil, err
	}

	p.ar.LogStats()

	var options []rainbow.Option
	options = append(options, optionsBTC...)
	options = append(options, optionsETH...)
	options = append(options, optionsSOL...)
	return options, nil
}

func (p *Provider) query(coin string) ([]instrument, error) {
	const baseURL = "https://deribit.com/api/v2/public/get_instruments?currency="
	const opts = "&expired=false&kind=option"
	url := baseURL + coin + opts
	log.Print("INF Deribit " + url)

	var result instrumentsResult
	err := p.ar.Get(coin, url, &result, maxBytesToRead)
	if err != nil {
		return nil, err
	}

	return filterTooFar(result.Result), nil
}

type instrumentsResult struct {
	Result []instrument `json:"result"`
}

type instrument struct {
	OptionType           string  `json:"option_type"`
	InstrumentName       string  `json:"instrument_name"`
	Kind                 string  `json:"kind"`
	SettlementPeriod     string  `json:"settlement_period"`
	QuoteCurrency        string  `json:"quote_currency"`
	BaseCurrency         string  `json:"base_currency"`
	MinTradeAmount       float64 `json:"min_trade_amount"`
	MakerCommission      float64 `json:"maker_commission"`
	Strike               float64 `json:"strike"`
	TickSize             float64 `json:"tick_size"`
	TakerCommission      float64 `json:"taker_commission"`
	ExpirationTimestamp  int64   `json:"expiration_timestamp"`
	CreationTimestamp    int64   `json:"creation_timestamp"`
	ContractSize         float64 `json:"contract_size"`
	BlockTradeCommission float64 `json:"block_trade_commission"`
	IsActive             bool    `json:"is_active"`
}

func filterTooFar(instruments []instrument) []instrument {
	expiries := rainbow.Expiries(time.Now(), Hour)

	filtered := make([]instrument, 0, len(instruments))

	for i := range instruments {
		seconds := instruments[i].ExpirationTimestamp / 1000
		ns := (instruments[i].ExpirationTimestamp % 1000) * 1000_000
		expiryTime := time.Unix(seconds, ns).UTC()
		// we should filter by taking what is available elsewhere and then
		// only fetch those
		if rainbow.IsExpiryAvailable(expiries, expiryTime) && isStrikeAvailable(&instruments[i]) {
			filtered = append(filtered, instruments[i])
		}
	}

	return filtered
}

// TODO change this quick and dirty way of filtering strikes from deribit.
func isStrikeAvailable(i *instrument) bool {
	ethStrike := []float64{500, 4000}
	btcStrike := []float64{9000, 40000}
	solStrike := []float64{10, 60}
	strikes := ethStrike

	if i.BaseCurrency == "BTC" {
		strikes = btcStrike
	} else if i.BaseCurrency == "SOL" {
		strikes = solStrike
	}

	return i.Strike >= strikes[0] && i.Strike <= strikes[1]
}

func (p *Provider) fillOptions(instruments []instrument, depth uint32) ([]rainbow.Option, error) {
	options := make([]rainbow.Option, 0, len(instruments))
	var lastError error

	baseURL := "https://www.deribit.com/api/v2/public/get_order_book?depth=" + strconv.Itoa(int(depth)) + "&instrument_name="

	for i := range instruments {
		var result orderBookResult
		url := baseURL + instruments[i].InstrumentName
		if err := p.ar.Get(instruments[i].InstrumentName, url, &result); err != nil {
			lastError = err
			log.Print("WRN Deribit book " + err.Error())
		}

		// API doc: https://docs.deribit.com/#public-get_index_price_names
		// ExpirationTimestamp = The time when the instrument will expire (milliseconds since the UNIX epoch)
		seconds := instruments[i].ExpirationTimestamp / 1000
		ns := (instruments[i].ExpirationTimestamp % 1000) * 1000_000
		expiryTime := time.Unix(seconds, ns).UTC()
		expiryStr := expiryTime.Format("2006-01-02 15:04:05")

		bids := normalizeOrders(result.Result.Bids, result.Result.IndexPrice)
		sort.Slice(bids, func(i, j int) bool {
			return bids[i].Price > bids[j].Price
		})

		asks := normalizeOrders(result.Result.Asks, result.Result.IndexPrice)
		sort.Slice(asks, func(i, j int) bool {
			return asks[i].Price < asks[j].Price
		})

		options = append(options, rainbow.Option{
			Name:          instruments[i].InstrumentName,
			Type:          strings.ToUpper(instruments[i].OptionType),
			Asset:         instruments[i].BaseCurrency,
			Expiry:        expiryStr,
			Strike:        instruments[i].Strike,
			ExchangeType:  "CEX",
			Chain:         "–",
			Layer:         "–",
			Provider:      "Deribit",
			QuoteCurrency: instruments[i].QuoteCurrency,
			Bid:           bids,
			Ask:           asks,
		})

		sleep(i) // rate limit the Deribit API
	}

	if len(options) == 0 {
		return nil, lastError
	}
	return options, nil
}

type orderBookResult struct {
	Result OrderBook `json:"result"`
}

type OrderBook struct {
	InstrumentName  string      `json:"instrument_name"`
	UnderlyingIndex string      `json:"underlying_index"`
	State           string      `json:"state"`
	Asks            [][]float64 `json:"asks"`
	Bids            [][]float64 `json:"bids"`
	Greeks          struct {
		Vega  float64 `json:"vega"`
		Theta float64 `json:"theta"`
		Rho   float64 `json:"rho"`
		Gamma float64 `json:"gamma"`
		Delta float64 `json:"delta"`
	} `json:"greeks"`
	Stats struct {
		Volume      float64 `json:"volume"`
		PriceChange float64 `json:"price_change"`
		Low         float64 `json:"low"`
		High        float64 `json:"high"`
	} `json:"stats"`
	UnderlyingPrice        float64 `json:"underlying_price"`
	MaxPrice               float64 `json:"max_price"`
	MarkPrice              float64 `json:"mark_price"`
	MarkIv                 float64 `json:"mark_iv"`
	LastPrice              float64 `json:"last_price"`
	InterestRate           float64 `json:"interest_rate"`
	MinPrice               float64 `json:"min_price"`
	IndexPrice             float64 `json:"index_price"`
	OpenInterest           float64 `json:"open_interest"`
	EstimatedDeliveryPrice float64 `json:"estimated_delivery_price"`
	ChangeID               int64   `json:"change_id"`
	SettlementPrice        float64 `json:"settlement_price"`
	BidIv                  float64 `json:"bid_iv"`
	BestBidPrice           float64 `json:"best_bid_price"`
	BestBidAmount          float64 `json:"best_bid_amount"`
	BestAskPrice           float64 `json:"best_ask_price"`
	BestAskAmount          float64 `json:"best_ask_amount"`
	Timestamp              int64   `json:"timestamp"`
	AskIv                  float64 `json:"ask_iv"`
}

// Prices are not in $ but in crypto so we need the coin (index) price to multiply
// and get the USD price.
func normalizeOrders(orders [][]float64, assetPrice float64) []rainbow.Order {
	// if there is no offer, send price=0.0, quant=0.0
	// hopefully we never an array of empty array
	if len(orders) == 0 {
		return []rainbow.Order{{Price: 0.0, Size: 0.0}}
	}

	offers := make([]rainbow.Order, 0, len(orders))
	for _, ord := range orders {
		offers = append(offers, rainbow.Order{Price: ord[0] * assetPrice, Size: ord[1]})
	}

	return offers
}

type AdaptiveRate struct {
	Name      string
	NextSleep time.Duration
	MinSleep  time.Duration
}

func NewAdaptiveRate(name string, d time.Duration) AdaptiveRate {
	ar := AdaptiveRate{
		Name:      name,
		NextSleep: d * factorInitialNextSleep,
		MinSleep:  d,
	}

	ar.LogStats()

	return ar
}

const (
	factorInitialNextSleep  = 2
	factorIncreaseMinSleep  = 32  // higher, the change is slower
	factorDecreaseMinSleep  = 512 // higher, the change is slower
	factorIncreaseNextSleep = 2   // higher, the change is faster
	factorDecreaseNextSleep = 8   // higher, the change is slower
	maxAlpha                = 16
	printDebug              = false
)

func (ar *AdaptiveRate) adjust(d time.Duration) {
	const fim = factorIncreaseMinSleep - 1
	const fin = factorIncreaseNextSleep - 1
	const fdn = factorDecreaseNextSleep - 1

	if d > ar.NextSleep {
		prevNext := ar.NextSleep
		prevMin := ar.MinSleep
		ar.NextSleep = (ar.NextSleep + fin*d) / factorIncreaseNextSleep
		ar.MinSleep = (d + fim*ar.MinSleep) / factorIncreaseMinSleep
		ar.logIncrease(prevMin, prevNext)
		return
	}

	// gap is used to detect stabilized sleep duration
	gap := ar.NextSleep - ar.MinSleep

	ar.NextSleep = (ar.MinSleep + fdn*ar.NextSleep) / factorDecreaseNextSleep

	// try to reduce slowly the "min sleep time"
	if reduce := ar.MinSleep / factorDecreaseMinSleep; gap < reduce {
		ar.MinSleep -= reduce
		ar.logDecrease(reduce)
	}
}

func (ar *AdaptiveRate) Get(symbol, url string, msg any, maxBytes ...int) error {
	var err error
	d := ar.NextSleep
	for try, status := 1, http.StatusTooManyRequests; (try < 88) && (status == http.StatusTooManyRequests); try++ {
		if try > 1 {
			previous := d
			alpha := int64(maxAlpha * ar.MinSleep / d)
			d *= time.Duration(try)
			d += time.Duration(alpha) * ar.MinSleep
			log.Printf("INF %s Get %s #%d sleep=%s (+%s) alpha=%d n=%s min=%s",
				ar.Name, symbol, try, d, d-previous, alpha, ar.NextSleep, ar.MinSleep)
		}
		time.Sleep(d)
		status, err = ar.get(symbol, url, msg, maxBytes...)
	}

	ar.adjust(d)

	return err
}

func (ar *AdaptiveRate) get(symbol, url string, msg any, maxBytes ...int) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return resp.StatusCode, fmt.Errorf("GET %s %s: %w", ar.Name, symbol, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		return resp.StatusCode, errors.New("Too Many Requests " + symbol)
	}

	if err = garcon.DecodeJSONResponse(resp, msg, maxBytes...); err != nil {
		return resp.StatusCode, fmt.Errorf("decode book %s: %w", symbol, err)
	}

	return resp.StatusCode, nil
}

func (ar *AdaptiveRate) LogStats() {
	log.Printf("INF %s Adjusted sleep durations: min=%s next=%s",
		ar.Name, ar.MinSleep, ar.NextSleep)
}

func (ar *AdaptiveRate) logIncrease(prevMin, prevNext time.Duration) {
	if printDebug {
		log.Printf("DBG %s Increase MinSleep=%s (+%s) next=%s (+%s)",
			ar.Name, ar.MinSleep, ar.MinSleep-prevMin, ar.NextSleep, ar.NextSleep-prevNext)
	}
}

func (ar *AdaptiveRate) logDecrease(reduce time.Duration) {
	if printDebug {
		log.Printf("DBG %s Decrease MinSleep=%s (-%s) next=%s",
			ar.Name, ar.MinSleep, reduce, ar.NextSleep)
	}
}
