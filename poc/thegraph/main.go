// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

package main

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spewerspew/spew"
	"github.com/teal-finance/rainbow/pkg/provider/thales"
	"github.com/teal-finance/rainbow/pkg/provider/zerox"
)

func main() {
	options := zerox.QueryTheGraph()
	spew.Dump(len(options))
	opt := thales.QueryMarket("0xd9e137e43fb366fb500553c2066b805616d9b0d4")
	spew.Dump(opt)

	/*
		const url = "https://api.thegraph.com/subgraphs/name/thales-markets/thales-optimism"
		//const url = "https://api.thegraph.com/subgraphs/name/thales-markets/thales-polygon"
		graphqlClient := graphql.NewClient(url, nil)
		resp, err := thales.AllLive(context.TODO(), graphqlClient) //, "0x5416c2ab11c7852ed9648aa006ee69d412c735c9")
		if err != nil {
			log.Print("ERR: ", err)
		}

		if resp == nil {
			log.Print("ERR: resp=nil")
			return
		}
		//spew.Dump(len(resp.Markets))
		/*
			var t int64 = 1654041600 //01-JUN-2022
			response, err := thales.Markets(context.TODO(), graphqlClient, 0, 0, t)
			if err != nil {
				log.Print("ERR: ", err)
			}

			if resp == nil {
				log.Print("ERR: resp=nil")
				return
			}
			spew.Dump(len(response.Markets))
			//name("0x4c554e4100000000000000000000000000000000000000000000000000000000")
			for _, m := range response.Markets {
				if m.Timestamp != 0 { //name(m.CurrencyKey) == "BTC" && m.ExpiryDate == int64(1672992000) {
					spew.Dump(m)
				}

			}
	*/
}

func name(c []byte) string {
	l := common.HexToHash(string(c))
	b := l.Bytes()
	b = bytes.Trim(b, "\x00")
	return string(b)
}
