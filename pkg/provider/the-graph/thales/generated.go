// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package thales

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// MarketMarket includes the requested fields of the GraphQL type Market.
type MarketMarket struct {
	Timestamp    int64  `json:"timestamp"`
	Creator      []byte `json:"creator"`
	CurrencyKey  []byte `json:"currencyKey"`
	StrikePrice  int64  `json:"strikePrice"`
	MaturityDate int64  `json:"maturityDate"`
	ExpiryDate   int64  `json:"expiryDate"`
	IsOpen       bool   `json:"isOpen"`
	PoolSize     int64  `json:"poolSize"`
	LongAddress  []byte `json:"longAddress"`
	ShortAddress []byte `json:"shortAddress"`
	Result       int    `json:"result"`
	CustomMarket bool   `json:"customMarket"`
	CustomOracle []byte `json:"customOracle"`
	FinalPrice   int64  `json:"finalPrice"`
}

// GetTimestamp returns MarketMarket.Timestamp, and is useful for accessing the field via an interface.
func (v *MarketMarket) GetTimestamp() int64 { return v.Timestamp }

// GetCreator returns MarketMarket.Creator, and is useful for accessing the field via an interface.
func (v *MarketMarket) GetCreator() []byte { return v.Creator }

// GetCurrencyKey returns MarketMarket.CurrencyKey, and is useful for accessing the field via an interface.
func (v *MarketMarket) GetCurrencyKey() []byte { return v.CurrencyKey }

// GetStrikePrice returns MarketMarket.StrikePrice, and is useful for accessing the field via an interface.
func (v *MarketMarket) GetStrikePrice() int64 { return v.StrikePrice }

// GetMaturityDate returns MarketMarket.MaturityDate, and is useful for accessing the field via an interface.
func (v *MarketMarket) GetMaturityDate() int64 { return v.MaturityDate }

// GetExpiryDate returns MarketMarket.ExpiryDate, and is useful for accessing the field via an interface.
func (v *MarketMarket) GetExpiryDate() int64 { return v.ExpiryDate }

// GetIsOpen returns MarketMarket.IsOpen, and is useful for accessing the field via an interface.
func (v *MarketMarket) GetIsOpen() bool { return v.IsOpen }

// GetPoolSize returns MarketMarket.PoolSize, and is useful for accessing the field via an interface.
func (v *MarketMarket) GetPoolSize() int64 { return v.PoolSize }

// GetLongAddress returns MarketMarket.LongAddress, and is useful for accessing the field via an interface.
func (v *MarketMarket) GetLongAddress() []byte { return v.LongAddress }

// GetShortAddress returns MarketMarket.ShortAddress, and is useful for accessing the field via an interface.
func (v *MarketMarket) GetShortAddress() []byte { return v.ShortAddress }

// GetResult returns MarketMarket.Result, and is useful for accessing the field via an interface.
func (v *MarketMarket) GetResult() int { return v.Result }

// GetCustomMarket returns MarketMarket.CustomMarket, and is useful for accessing the field via an interface.
func (v *MarketMarket) GetCustomMarket() bool { return v.CustomMarket }

// GetCustomOracle returns MarketMarket.CustomOracle, and is useful for accessing the field via an interface.
func (v *MarketMarket) GetCustomOracle() []byte { return v.CustomOracle }

// GetFinalPrice returns MarketMarket.FinalPrice, and is useful for accessing the field via an interface.
func (v *MarketMarket) GetFinalPrice() int64 { return v.FinalPrice }

// MarketResponse is returned by Market on success.
type MarketResponse struct {
	Market MarketMarket `json:"market"`
}

// GetMarket returns MarketResponse.Market, and is useful for accessing the field via an interface.
func (v *MarketResponse) GetMarket() MarketMarket { return v.Market }

// MarketsMarketsMarket includes the requested fields of the GraphQL type Market.
type MarketsMarketsMarket struct {
	Id           string `json:"id"`
	Timestamp    int64  `json:"timestamp"`
	Creator      []byte `json:"creator"`
	CurrencyKey  []byte `json:"currencyKey"`
	StrikePrice  int64  `json:"strikePrice"`
	MaturityDate int64  `json:"maturityDate"`
	ExpiryDate   int64  `json:"expiryDate"`
	IsOpen       bool   `json:"isOpen"`
	PoolSize     int64  `json:"poolSize"`
	LongAddress  []byte `json:"longAddress"`
	ShortAddress []byte `json:"shortAddress"`
	Result       int    `json:"result"`
	CustomMarket bool   `json:"customMarket"`
	CustomOracle []byte `json:"customOracle"`
	FinalPrice   int64  `json:"finalPrice"`
}

// GetId returns MarketsMarketsMarket.Id, and is useful for accessing the field via an interface.
func (v *MarketsMarketsMarket) GetId() string { return v.Id }

// GetTimestamp returns MarketsMarketsMarket.Timestamp, and is useful for accessing the field via an interface.
func (v *MarketsMarketsMarket) GetTimestamp() int64 { return v.Timestamp }

// GetCreator returns MarketsMarketsMarket.Creator, and is useful for accessing the field via an interface.
func (v *MarketsMarketsMarket) GetCreator() []byte { return v.Creator }

// GetCurrencyKey returns MarketsMarketsMarket.CurrencyKey, and is useful for accessing the field via an interface.
func (v *MarketsMarketsMarket) GetCurrencyKey() []byte { return v.CurrencyKey }

// GetStrikePrice returns MarketsMarketsMarket.StrikePrice, and is useful for accessing the field via an interface.
func (v *MarketsMarketsMarket) GetStrikePrice() int64 { return v.StrikePrice }

// GetMaturityDate returns MarketsMarketsMarket.MaturityDate, and is useful for accessing the field via an interface.
func (v *MarketsMarketsMarket) GetMaturityDate() int64 { return v.MaturityDate }

// GetExpiryDate returns MarketsMarketsMarket.ExpiryDate, and is useful for accessing the field via an interface.
func (v *MarketsMarketsMarket) GetExpiryDate() int64 { return v.ExpiryDate }

// GetIsOpen returns MarketsMarketsMarket.IsOpen, and is useful for accessing the field via an interface.
func (v *MarketsMarketsMarket) GetIsOpen() bool { return v.IsOpen }

// GetPoolSize returns MarketsMarketsMarket.PoolSize, and is useful for accessing the field via an interface.
func (v *MarketsMarketsMarket) GetPoolSize() int64 { return v.PoolSize }

// GetLongAddress returns MarketsMarketsMarket.LongAddress, and is useful for accessing the field via an interface.
func (v *MarketsMarketsMarket) GetLongAddress() []byte { return v.LongAddress }

// GetShortAddress returns MarketsMarketsMarket.ShortAddress, and is useful for accessing the field via an interface.
func (v *MarketsMarketsMarket) GetShortAddress() []byte { return v.ShortAddress }

// GetResult returns MarketsMarketsMarket.Result, and is useful for accessing the field via an interface.
func (v *MarketsMarketsMarket) GetResult() int { return v.Result }

// GetCustomMarket returns MarketsMarketsMarket.CustomMarket, and is useful for accessing the field via an interface.
func (v *MarketsMarketsMarket) GetCustomMarket() bool { return v.CustomMarket }

// GetCustomOracle returns MarketsMarketsMarket.CustomOracle, and is useful for accessing the field via an interface.
func (v *MarketsMarketsMarket) GetCustomOracle() []byte { return v.CustomOracle }

// GetFinalPrice returns MarketsMarketsMarket.FinalPrice, and is useful for accessing the field via an interface.
func (v *MarketsMarketsMarket) GetFinalPrice() int64 { return v.FinalPrice }

// MarketsResponse is returned by Markets on success.
type MarketsResponse struct {
	Markets []MarketsMarketsMarket `json:"markets"`
}

// GetMarkets returns MarketsResponse.Markets, and is useful for accessing the field via an interface.
func (v *MarketsResponse) GetMarkets() []MarketsMarketsMarket { return v.Markets }

// RangedMarketRangedMarket includes the requested fields of the GraphQL type RangedMarket.
type RangedMarketRangedMarket struct {
	Timestamp    int64                               `json:"timestamp"`
	CurrencyKey  []byte                              `json:"currencyKey"`
	MaturityDate int64                               `json:"maturityDate"`
	LeftPrice    int64                               `json:"leftPrice"`
	RightPrice   int64                               `json:"rightPrice"`
	InAddress    []byte                              `json:"inAddress"`
	OutAddress   []byte                              `json:"outAddress"`
	RightMarket  RangedMarketRangedMarketRightMarket `json:"rightMarket"`
	LeftMarket   RangedMarketRangedMarketLeftMarket  `json:"leftMarket"`
	IsOpen       bool                                `json:"isOpen"`
	Result       int                                 `json:"result"`
	FinalPrice   int64                               `json:"finalPrice"`
}

// GetTimestamp returns RangedMarketRangedMarket.Timestamp, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarket) GetTimestamp() int64 { return v.Timestamp }

// GetCurrencyKey returns RangedMarketRangedMarket.CurrencyKey, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarket) GetCurrencyKey() []byte { return v.CurrencyKey }

// GetMaturityDate returns RangedMarketRangedMarket.MaturityDate, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarket) GetMaturityDate() int64 { return v.MaturityDate }

// GetLeftPrice returns RangedMarketRangedMarket.LeftPrice, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarket) GetLeftPrice() int64 { return v.LeftPrice }

// GetRightPrice returns RangedMarketRangedMarket.RightPrice, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarket) GetRightPrice() int64 { return v.RightPrice }

// GetInAddress returns RangedMarketRangedMarket.InAddress, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarket) GetInAddress() []byte { return v.InAddress }

// GetOutAddress returns RangedMarketRangedMarket.OutAddress, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarket) GetOutAddress() []byte { return v.OutAddress }

// GetRightMarket returns RangedMarketRangedMarket.RightMarket, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarket) GetRightMarket() RangedMarketRangedMarketRightMarket {
	return v.RightMarket
}

// GetLeftMarket returns RangedMarketRangedMarket.LeftMarket, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarket) GetLeftMarket() RangedMarketRangedMarketLeftMarket {
	return v.LeftMarket
}

// GetIsOpen returns RangedMarketRangedMarket.IsOpen, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarket) GetIsOpen() bool { return v.IsOpen }

// GetResult returns RangedMarketRangedMarket.Result, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarket) GetResult() int { return v.Result }

// GetFinalPrice returns RangedMarketRangedMarket.FinalPrice, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarket) GetFinalPrice() int64 { return v.FinalPrice }

// RangedMarketRangedMarketLeftMarket includes the requested fields of the GraphQL type Market.
type RangedMarketRangedMarketLeftMarket struct {
	Id           string `json:"id"`
	Timestamp    int64  `json:"timestamp"`
	Creator      []byte `json:"creator"`
	CurrencyKey  []byte `json:"currencyKey"`
	StrikePrice  int64  `json:"strikePrice"`
	MaturityDate int64  `json:"maturityDate"`
	ExpiryDate   int64  `json:"expiryDate"`
	IsOpen       bool   `json:"isOpen"`
	PoolSize     int64  `json:"poolSize"`
	LongAddress  []byte `json:"longAddress"`
	ShortAddress []byte `json:"shortAddress"`
	Result       int    `json:"result"`
	CustomMarket bool   `json:"customMarket"`
	CustomOracle []byte `json:"customOracle"`
	FinalPrice   int64  `json:"finalPrice"`
}

// GetId returns RangedMarketRangedMarketLeftMarket.Id, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketLeftMarket) GetId() string { return v.Id }

// GetTimestamp returns RangedMarketRangedMarketLeftMarket.Timestamp, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketLeftMarket) GetTimestamp() int64 { return v.Timestamp }

// GetCreator returns RangedMarketRangedMarketLeftMarket.Creator, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketLeftMarket) GetCreator() []byte { return v.Creator }

// GetCurrencyKey returns RangedMarketRangedMarketLeftMarket.CurrencyKey, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketLeftMarket) GetCurrencyKey() []byte { return v.CurrencyKey }

// GetStrikePrice returns RangedMarketRangedMarketLeftMarket.StrikePrice, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketLeftMarket) GetStrikePrice() int64 { return v.StrikePrice }

// GetMaturityDate returns RangedMarketRangedMarketLeftMarket.MaturityDate, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketLeftMarket) GetMaturityDate() int64 { return v.MaturityDate }

// GetExpiryDate returns RangedMarketRangedMarketLeftMarket.ExpiryDate, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketLeftMarket) GetExpiryDate() int64 { return v.ExpiryDate }

// GetIsOpen returns RangedMarketRangedMarketLeftMarket.IsOpen, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketLeftMarket) GetIsOpen() bool { return v.IsOpen }

// GetPoolSize returns RangedMarketRangedMarketLeftMarket.PoolSize, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketLeftMarket) GetPoolSize() int64 { return v.PoolSize }

// GetLongAddress returns RangedMarketRangedMarketLeftMarket.LongAddress, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketLeftMarket) GetLongAddress() []byte { return v.LongAddress }

// GetShortAddress returns RangedMarketRangedMarketLeftMarket.ShortAddress, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketLeftMarket) GetShortAddress() []byte { return v.ShortAddress }

// GetResult returns RangedMarketRangedMarketLeftMarket.Result, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketLeftMarket) GetResult() int { return v.Result }

// GetCustomMarket returns RangedMarketRangedMarketLeftMarket.CustomMarket, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketLeftMarket) GetCustomMarket() bool { return v.CustomMarket }

// GetCustomOracle returns RangedMarketRangedMarketLeftMarket.CustomOracle, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketLeftMarket) GetCustomOracle() []byte { return v.CustomOracle }

// GetFinalPrice returns RangedMarketRangedMarketLeftMarket.FinalPrice, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketLeftMarket) GetFinalPrice() int64 { return v.FinalPrice }

// RangedMarketRangedMarketRightMarket includes the requested fields of the GraphQL type Market.
type RangedMarketRangedMarketRightMarket struct {
	Id           string `json:"id"`
	Timestamp    int64  `json:"timestamp"`
	Creator      []byte `json:"creator"`
	CurrencyKey  []byte `json:"currencyKey"`
	StrikePrice  int64  `json:"strikePrice"`
	MaturityDate int64  `json:"maturityDate"`
	ExpiryDate   int64  `json:"expiryDate"`
	IsOpen       bool   `json:"isOpen"`
	PoolSize     int64  `json:"poolSize"`
	LongAddress  []byte `json:"longAddress"`
	ShortAddress []byte `json:"shortAddress"`
	Result       int    `json:"result"`
	CustomMarket bool   `json:"customMarket"`
	CustomOracle []byte `json:"customOracle"`
	FinalPrice   int64  `json:"finalPrice"`
}

// GetId returns RangedMarketRangedMarketRightMarket.Id, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketRightMarket) GetId() string { return v.Id }

// GetTimestamp returns RangedMarketRangedMarketRightMarket.Timestamp, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketRightMarket) GetTimestamp() int64 { return v.Timestamp }

// GetCreator returns RangedMarketRangedMarketRightMarket.Creator, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketRightMarket) GetCreator() []byte { return v.Creator }

// GetCurrencyKey returns RangedMarketRangedMarketRightMarket.CurrencyKey, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketRightMarket) GetCurrencyKey() []byte { return v.CurrencyKey }

// GetStrikePrice returns RangedMarketRangedMarketRightMarket.StrikePrice, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketRightMarket) GetStrikePrice() int64 { return v.StrikePrice }

// GetMaturityDate returns RangedMarketRangedMarketRightMarket.MaturityDate, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketRightMarket) GetMaturityDate() int64 { return v.MaturityDate }

// GetExpiryDate returns RangedMarketRangedMarketRightMarket.ExpiryDate, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketRightMarket) GetExpiryDate() int64 { return v.ExpiryDate }

// GetIsOpen returns RangedMarketRangedMarketRightMarket.IsOpen, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketRightMarket) GetIsOpen() bool { return v.IsOpen }

// GetPoolSize returns RangedMarketRangedMarketRightMarket.PoolSize, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketRightMarket) GetPoolSize() int64 { return v.PoolSize }

// GetLongAddress returns RangedMarketRangedMarketRightMarket.LongAddress, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketRightMarket) GetLongAddress() []byte { return v.LongAddress }

// GetShortAddress returns RangedMarketRangedMarketRightMarket.ShortAddress, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketRightMarket) GetShortAddress() []byte { return v.ShortAddress }

// GetResult returns RangedMarketRangedMarketRightMarket.Result, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketRightMarket) GetResult() int { return v.Result }

// GetCustomMarket returns RangedMarketRangedMarketRightMarket.CustomMarket, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketRightMarket) GetCustomMarket() bool { return v.CustomMarket }

// GetCustomOracle returns RangedMarketRangedMarketRightMarket.CustomOracle, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketRightMarket) GetCustomOracle() []byte { return v.CustomOracle }

// GetFinalPrice returns RangedMarketRangedMarketRightMarket.FinalPrice, and is useful for accessing the field via an interface.
func (v *RangedMarketRangedMarketRightMarket) GetFinalPrice() int64 { return v.FinalPrice }

// RangedMarketResponse is returned by RangedMarket on success.
type RangedMarketResponse struct {
	RangedMarket RangedMarketRangedMarket `json:"rangedMarket"`
}

// GetRangedMarket returns RangedMarketResponse.RangedMarket, and is useful for accessing the field via an interface.
func (v *RangedMarketResponse) GetRangedMarket() RangedMarketRangedMarket { return v.RangedMarket }

// RangedMarketsRangedMarketsRangedMarket includes the requested fields of the GraphQL type RangedMarket.
type RangedMarketsRangedMarketsRangedMarket struct {
	Id           string                                            `json:"id"`
	Timestamp    int64                                             `json:"timestamp"`
	CurrencyKey  []byte                                            `json:"currencyKey"`
	MaturityDate int64                                             `json:"maturityDate"`
	LeftPrice    int64                                             `json:"leftPrice"`
	RightPrice   int64                                             `json:"rightPrice"`
	InAddress    []byte                                            `json:"inAddress"`
	OutAddress   []byte                                            `json:"outAddress"`
	RightMarket  RangedMarketsRangedMarketsRangedMarketRightMarket `json:"rightMarket"`
	LeftMarket   RangedMarketsRangedMarketsRangedMarketLeftMarket  `json:"leftMarket"`
	IsOpen       bool                                              `json:"isOpen"`
	Result       int                                               `json:"result"`
	FinalPrice   int64                                             `json:"finalPrice"`
}

// GetId returns RangedMarketsRangedMarketsRangedMarket.Id, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarket) GetId() string { return v.Id }

// GetTimestamp returns RangedMarketsRangedMarketsRangedMarket.Timestamp, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarket) GetTimestamp() int64 { return v.Timestamp }

// GetCurrencyKey returns RangedMarketsRangedMarketsRangedMarket.CurrencyKey, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarket) GetCurrencyKey() []byte { return v.CurrencyKey }

// GetMaturityDate returns RangedMarketsRangedMarketsRangedMarket.MaturityDate, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarket) GetMaturityDate() int64 { return v.MaturityDate }

// GetLeftPrice returns RangedMarketsRangedMarketsRangedMarket.LeftPrice, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarket) GetLeftPrice() int64 { return v.LeftPrice }

// GetRightPrice returns RangedMarketsRangedMarketsRangedMarket.RightPrice, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarket) GetRightPrice() int64 { return v.RightPrice }

// GetInAddress returns RangedMarketsRangedMarketsRangedMarket.InAddress, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarket) GetInAddress() []byte { return v.InAddress }

// GetOutAddress returns RangedMarketsRangedMarketsRangedMarket.OutAddress, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarket) GetOutAddress() []byte { return v.OutAddress }

// GetRightMarket returns RangedMarketsRangedMarketsRangedMarket.RightMarket, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarket) GetRightMarket() RangedMarketsRangedMarketsRangedMarketRightMarket {
	return v.RightMarket
}

// GetLeftMarket returns RangedMarketsRangedMarketsRangedMarket.LeftMarket, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarket) GetLeftMarket() RangedMarketsRangedMarketsRangedMarketLeftMarket {
	return v.LeftMarket
}

// GetIsOpen returns RangedMarketsRangedMarketsRangedMarket.IsOpen, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarket) GetIsOpen() bool { return v.IsOpen }

// GetResult returns RangedMarketsRangedMarketsRangedMarket.Result, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarket) GetResult() int { return v.Result }

// GetFinalPrice returns RangedMarketsRangedMarketsRangedMarket.FinalPrice, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarket) GetFinalPrice() int64 { return v.FinalPrice }

// RangedMarketsRangedMarketsRangedMarketLeftMarket includes the requested fields of the GraphQL type Market.
type RangedMarketsRangedMarketsRangedMarketLeftMarket struct {
	Id           string `json:"id"`
	Timestamp    int64  `json:"timestamp"`
	Creator      []byte `json:"creator"`
	CurrencyKey  []byte `json:"currencyKey"`
	StrikePrice  int64  `json:"strikePrice"`
	MaturityDate int64  `json:"maturityDate"`
	ExpiryDate   int64  `json:"expiryDate"`
	IsOpen       bool   `json:"isOpen"`
	PoolSize     int64  `json:"poolSize"`
	LongAddress  []byte `json:"longAddress"`
	ShortAddress []byte `json:"shortAddress"`
	Result       int    `json:"result"`
	CustomMarket bool   `json:"customMarket"`
	CustomOracle []byte `json:"customOracle"`
	FinalPrice   int64  `json:"finalPrice"`
}

// GetId returns RangedMarketsRangedMarketsRangedMarketLeftMarket.Id, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketLeftMarket) GetId() string { return v.Id }

// GetTimestamp returns RangedMarketsRangedMarketsRangedMarketLeftMarket.Timestamp, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketLeftMarket) GetTimestamp() int64 { return v.Timestamp }

// GetCreator returns RangedMarketsRangedMarketsRangedMarketLeftMarket.Creator, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketLeftMarket) GetCreator() []byte { return v.Creator }

// GetCurrencyKey returns RangedMarketsRangedMarketsRangedMarketLeftMarket.CurrencyKey, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketLeftMarket) GetCurrencyKey() []byte {
	return v.CurrencyKey
}

// GetStrikePrice returns RangedMarketsRangedMarketsRangedMarketLeftMarket.StrikePrice, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketLeftMarket) GetStrikePrice() int64 {
	return v.StrikePrice
}

// GetMaturityDate returns RangedMarketsRangedMarketsRangedMarketLeftMarket.MaturityDate, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketLeftMarket) GetMaturityDate() int64 {
	return v.MaturityDate
}

// GetExpiryDate returns RangedMarketsRangedMarketsRangedMarketLeftMarket.ExpiryDate, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketLeftMarket) GetExpiryDate() int64 { return v.ExpiryDate }

// GetIsOpen returns RangedMarketsRangedMarketsRangedMarketLeftMarket.IsOpen, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketLeftMarket) GetIsOpen() bool { return v.IsOpen }

// GetPoolSize returns RangedMarketsRangedMarketsRangedMarketLeftMarket.PoolSize, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketLeftMarket) GetPoolSize() int64 { return v.PoolSize }

// GetLongAddress returns RangedMarketsRangedMarketsRangedMarketLeftMarket.LongAddress, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketLeftMarket) GetLongAddress() []byte {
	return v.LongAddress
}

// GetShortAddress returns RangedMarketsRangedMarketsRangedMarketLeftMarket.ShortAddress, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketLeftMarket) GetShortAddress() []byte {
	return v.ShortAddress
}

// GetResult returns RangedMarketsRangedMarketsRangedMarketLeftMarket.Result, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketLeftMarket) GetResult() int { return v.Result }

// GetCustomMarket returns RangedMarketsRangedMarketsRangedMarketLeftMarket.CustomMarket, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketLeftMarket) GetCustomMarket() bool {
	return v.CustomMarket
}

// GetCustomOracle returns RangedMarketsRangedMarketsRangedMarketLeftMarket.CustomOracle, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketLeftMarket) GetCustomOracle() []byte {
	return v.CustomOracle
}

// GetFinalPrice returns RangedMarketsRangedMarketsRangedMarketLeftMarket.FinalPrice, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketLeftMarket) GetFinalPrice() int64 { return v.FinalPrice }

// RangedMarketsRangedMarketsRangedMarketRightMarket includes the requested fields of the GraphQL type Market.
type RangedMarketsRangedMarketsRangedMarketRightMarket struct {
	Id           string `json:"id"`
	Timestamp    int64  `json:"timestamp"`
	Creator      []byte `json:"creator"`
	CurrencyKey  []byte `json:"currencyKey"`
	StrikePrice  int64  `json:"strikePrice"`
	MaturityDate int64  `json:"maturityDate"`
	ExpiryDate   int64  `json:"expiryDate"`
	IsOpen       bool   `json:"isOpen"`
	PoolSize     int64  `json:"poolSize"`
	LongAddress  []byte `json:"longAddress"`
	ShortAddress []byte `json:"shortAddress"`
	Result       int    `json:"result"`
	CustomMarket bool   `json:"customMarket"`
	CustomOracle []byte `json:"customOracle"`
	FinalPrice   int64  `json:"finalPrice"`
}

// GetId returns RangedMarketsRangedMarketsRangedMarketRightMarket.Id, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketRightMarket) GetId() string { return v.Id }

// GetTimestamp returns RangedMarketsRangedMarketsRangedMarketRightMarket.Timestamp, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketRightMarket) GetTimestamp() int64 { return v.Timestamp }

// GetCreator returns RangedMarketsRangedMarketsRangedMarketRightMarket.Creator, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketRightMarket) GetCreator() []byte { return v.Creator }

// GetCurrencyKey returns RangedMarketsRangedMarketsRangedMarketRightMarket.CurrencyKey, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketRightMarket) GetCurrencyKey() []byte {
	return v.CurrencyKey
}

// GetStrikePrice returns RangedMarketsRangedMarketsRangedMarketRightMarket.StrikePrice, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketRightMarket) GetStrikePrice() int64 {
	return v.StrikePrice
}

// GetMaturityDate returns RangedMarketsRangedMarketsRangedMarketRightMarket.MaturityDate, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketRightMarket) GetMaturityDate() int64 {
	return v.MaturityDate
}

// GetExpiryDate returns RangedMarketsRangedMarketsRangedMarketRightMarket.ExpiryDate, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketRightMarket) GetExpiryDate() int64 {
	return v.ExpiryDate
}

// GetIsOpen returns RangedMarketsRangedMarketsRangedMarketRightMarket.IsOpen, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketRightMarket) GetIsOpen() bool { return v.IsOpen }

// GetPoolSize returns RangedMarketsRangedMarketsRangedMarketRightMarket.PoolSize, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketRightMarket) GetPoolSize() int64 { return v.PoolSize }

// GetLongAddress returns RangedMarketsRangedMarketsRangedMarketRightMarket.LongAddress, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketRightMarket) GetLongAddress() []byte {
	return v.LongAddress
}

// GetShortAddress returns RangedMarketsRangedMarketsRangedMarketRightMarket.ShortAddress, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketRightMarket) GetShortAddress() []byte {
	return v.ShortAddress
}

// GetResult returns RangedMarketsRangedMarketsRangedMarketRightMarket.Result, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketRightMarket) GetResult() int { return v.Result }

// GetCustomMarket returns RangedMarketsRangedMarketsRangedMarketRightMarket.CustomMarket, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketRightMarket) GetCustomMarket() bool {
	return v.CustomMarket
}

// GetCustomOracle returns RangedMarketsRangedMarketsRangedMarketRightMarket.CustomOracle, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketRightMarket) GetCustomOracle() []byte {
	return v.CustomOracle
}

// GetFinalPrice returns RangedMarketsRangedMarketsRangedMarketRightMarket.FinalPrice, and is useful for accessing the field via an interface.
func (v *RangedMarketsRangedMarketsRangedMarketRightMarket) GetFinalPrice() int64 {
	return v.FinalPrice
}

// RangedMarketsResponse is returned by RangedMarkets on success.
type RangedMarketsResponse struct {
	RangedMarkets []RangedMarketsRangedMarketsRangedMarket `json:"rangedMarkets"`
}

// GetRangedMarkets returns RangedMarketsResponse.RangedMarkets, and is useful for accessing the field via an interface.
func (v *RangedMarketsResponse) GetRangedMarkets() []RangedMarketsRangedMarketsRangedMarket {
	return v.RangedMarkets
}

// __MarketInput is used internally by genqlient
type __MarketInput struct {
	Id string `json:"id"`
}

// GetId returns __MarketInput.Id, and is useful for accessing the field via an interface.
func (v *__MarketInput) GetId() string { return v.Id }

// __MarketsInput is used internally by genqlient
type __MarketsInput struct {
	Skip  int   `json:"skip"`
	First int   `json:"first"`
	T     int64 `json:"t"`
}

// GetSkip returns __MarketsInput.Skip, and is useful for accessing the field via an interface.
func (v *__MarketsInput) GetSkip() int { return v.Skip }

// GetFirst returns __MarketsInput.First, and is useful for accessing the field via an interface.
func (v *__MarketsInput) GetFirst() int { return v.First }

// GetT returns __MarketsInput.T, and is useful for accessing the field via an interface.
func (v *__MarketsInput) GetT() int64 { return v.T }

// __RangedMarketInput is used internally by genqlient
type __RangedMarketInput struct {
	Id string `json:"id"`
}

// GetId returns __RangedMarketInput.Id, and is useful for accessing the field via an interface.
func (v *__RangedMarketInput) GetId() string { return v.Id }

// __RangedMarketsInput is used internally by genqlient
type __RangedMarketsInput struct {
	Skip  int   `json:"skip"`
	First int   `json:"first"`
	T     int64 `json:"t"`
}

// GetSkip returns __RangedMarketsInput.Skip, and is useful for accessing the field via an interface.
func (v *__RangedMarketsInput) GetSkip() int { return v.Skip }

// GetFirst returns __RangedMarketsInput.First, and is useful for accessing the field via an interface.
func (v *__RangedMarketsInput) GetFirst() int { return v.First }

// GetT returns __RangedMarketsInput.T, and is useful for accessing the field via an interface.
func (v *__RangedMarketsInput) GetT() int64 { return v.T }

func Market(
	ctx context.Context,
	client graphql.Client,
	id string,
) (*MarketResponse, error) {
	__input := __MarketInput{
		Id: id,
	}
	var err error

	var retval MarketResponse
	err = client.MakeRequest(
		ctx,
		"Market",
		`
query Market ($id: ID!) {
	market(id: $id) {
		timestamp
		creator
		currencyKey
		strikePrice
		maturityDate
		expiryDate
		isOpen
		poolSize
		longAddress
		shortAddress
		result
		customMarket
		customOracle
		finalPrice
	}
}
`,
		&retval,
		&__input,
	)
	return &retval, err
}

func Markets(
	ctx context.Context,
	client graphql.Client,
	skip int,
	first int,
	t int64,
) (*MarketsResponse, error) {
	__input := __MarketsInput{
		Skip:  skip,
		First: first,
		T:     t,
	}
	var err error

	var retval MarketsResponse
	err = client.MakeRequest(
		ctx,
		"Markets",
		`
query Markets ($skip: Int, $first: Int, $t: BigInt) {
	markets(skip: $skip, first: $first, orderBy: timestamp, orderDirection: desc, where: {timestamp_gt:$t}) {
		id
		timestamp
		creator
		currencyKey
		strikePrice
		maturityDate
		expiryDate
		isOpen
		poolSize
		longAddress
		shortAddress
		result
		customMarket
		customOracle
		finalPrice
	}
}
`,
		&retval,
		&__input,
	)
	return &retval, err
}

func RangedMarket(
	ctx context.Context,
	client graphql.Client,
	id string,
) (*RangedMarketResponse, error) {
	__input := __RangedMarketInput{
		Id: id,
	}
	var err error

	var retval RangedMarketResponse
	err = client.MakeRequest(
		ctx,
		"RangedMarket",
		`
query RangedMarket ($id: ID!) {
	rangedMarket(id: $id) {
		timestamp
		currencyKey
		maturityDate
		leftPrice
		rightPrice
		inAddress
		outAddress
		rightMarket {
			id
			timestamp
			creator
			currencyKey
			strikePrice
			maturityDate
			expiryDate
			isOpen
			poolSize
			longAddress
			shortAddress
			result
			customMarket
			customOracle
			finalPrice
		}
		leftMarket {
			id
			timestamp
			creator
			currencyKey
			strikePrice
			maturityDate
			expiryDate
			isOpen
			poolSize
			longAddress
			shortAddress
			result
			customMarket
			customOracle
			finalPrice
		}
		isOpen
		result
		finalPrice
	}
}
`,
		&retval,
		&__input,
	)
	return &retval, err
}

func RangedMarkets(
	ctx context.Context,
	client graphql.Client,
	skip int,
	first int,
	t int64,
) (*RangedMarketsResponse, error) {
	__input := __RangedMarketsInput{
		Skip:  skip,
		First: first,
		T:     t,
	}
	var err error

	var retval RangedMarketsResponse
	err = client.MakeRequest(
		ctx,
		"RangedMarkets",
		`
query RangedMarkets ($skip: Int, $first: Int, $t: BigInt) {
	rangedMarkets(skip: $skip, first: $first, orderBy: timestamp, orderDirection: desc, where: {timestamp_gt:$t}) {
		id
		timestamp
		currencyKey
		maturityDate
		leftPrice
		rightPrice
		inAddress
		outAddress
		rightMarket {
			id
			timestamp
			creator
			currencyKey
			strikePrice
			maturityDate
			expiryDate
			isOpen
			poolSize
			longAddress
			shortAddress
			result
			customMarket
			customOracle
			finalPrice
		}
		leftMarket {
			id
			timestamp
			creator
			currencyKey
			strikePrice
			maturityDate
			expiryDate
			isOpen
			poolSize
			longAddress
			shortAddress
			result
			customMarket
			customOracle
			finalPrice
		}
		isOpen
		result
		finalPrice
	}
}
`,
		&retval,
		&__input,
	)
	return &retval, err
}
