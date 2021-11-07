// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package rainbow

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	c *Cache
}

func (h *handler) router() http.Handler {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/options", h.getOptions)
		r.Get("/options/cp", h.getCPFormat)
	})

	return r
}

func (h handler) getOptions(w http.ResponseWriter, r *http.Request) {
	options, err := h.c.Options()
	if err != nil {
		log.Print("ERROR getOptions ", err)
		http.Error(w, "No Content", http.StatusNoContent)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(options); err != nil {
		log.Print("ERROR getOptions ", err)
		http.Error(w, "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)

		return
	}
}

func (h handler) getCPFormat(w http.ResponseWriter, r *http.Request) {
	cp, err := h.c.CPFormat()
	if err != nil {
		log.Print("ERROR getCPFormat ", err)
		http.Error(w, "No Content", http.StatusNoContent)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(cp); err != nil {
		log.Print("ERROR getCPFormat ", err)
		http.Error(w, "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)

		return
	}
}

type CPFormat struct {
	Rows []Row `json:"rows"`
}

type Row struct {
	Asset    string `json:"asset"`
	Expiry   string `json:"expiry"`
	Provider string `json:"provider"`

	Call OptionIndicators `json:"call"`

	Strike float64 `json:"strike"`

	Put OptionIndicators `json:"put"`
}

type OptionIndicators struct {
	Bid SimpleOrder `json:"bid"`
	Ask SimpleOrder `json:"ask"`
}

type SimpleOrder struct {
	Price float64 `json:"px"`
	Size  float64 `json:"size"`
}

func buildCPFormat(options []Option) CPFormat {
	rows := make([]Row, 0, len(options)/2)

	for asset, optionsSameAsset := range groupByAsset(options) {
		for expiry, optionsSameExpiry := range groupByExpiry(optionsSameAsset) {
			for strike, optionsSameStrike := range groupByStrike(optionsSameExpiry) {
				for provider, optionsSameProvider := range groupByProvider(optionsSameStrike) {
					r := Row{
						Asset:    asset,
						Expiry:   expiry,
						Provider: provider,
						Strike:   strike,
					}

					for _, o := range optionsSameProvider {
						if o.Type == "PUT" {
							r.Put = newOptionIndicators(o)
						} else {
							r.Call = newOptionIndicators(o)
						}
					}

					rows = append(rows, r)
				}
			}
		}
	}

	return CPFormat{Rows: rows}
}

func groupByAsset(options []Option) (assetToOptions map[string][]Option) {
	assetToOptions = map[string][]Option{}

	for _, o := range options {
		slice, ok := assetToOptions[o.Asset]
		if ok {
			assetToOptions[o.Asset] = append(slice, o)
		} else {
			assetToOptions[o.Asset] = []Option{o}
		}
	}

	return assetToOptions
}

func groupByExpiry(options []Option) (expiryToOptions map[string][]Option) {
	expiryToOptions = map[string][]Option{}

	for _, o := range options {
		slice, ok := expiryToOptions[o.Expiry]
		if ok {
			expiryToOptions[o.Expiry] = append(slice, o)
		} else {
			expiryToOptions[o.Expiry] = []Option{o}
		}
	}

	return expiryToOptions
}

func groupByStrike(options []Option) (strikeToOptions map[float64][]Option) {
	strikeToOptions = map[float64][]Option{}

	for _, o := range options {
		slice, ok := strikeToOptions[o.Strike]
		if ok {
			strikeToOptions[o.Strike] = append(slice, o)
		} else {
			strikeToOptions[o.Strike] = []Option{o}
		}
	}

	return strikeToOptions
}

func groupByProvider(options []Option) (providerToOptions map[string][]Option) {
	providerToOptions = map[string][]Option{}

	for _, o := range options {
		slice, ok := providerToOptions[o.Provider]
		if ok {
			providerToOptions[o.Provider] = append(slice, o)
		} else {
			providerToOptions[o.Provider] = []Option{o}
		}
	}

	return providerToOptions
}

func newOptionIndicators(o Option) OptionIndicators {
	oi := OptionIndicators{
		Bid: SimpleOrder{Price: 0, Size: 0},
		Ask: SimpleOrder{Price: 0, Size: 0},
	}

	if len(o.Bid) > 0 {
		oi.Bid.Price = o.Bid[0].Price
		oi.Bid.Size = o.Bid[0].Quantity
	}

	if len(o.Ask) > 0 {
		oi.Ask.Price = o.Ask[0].Price
		oi.Ask.Size = o.Ask[0].Quantity
	}

	return oi
}