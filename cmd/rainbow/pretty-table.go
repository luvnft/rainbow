// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gookit/color"
	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/teal-finance/rainbow"
	"github.com/teal-finance/rainbow/pkg/all"
)

func printPrettyTable() {
	options, err := all.OptionsFromAllProviders()
	if err != nil {
		log.Print("ERROR: ", err)
		return
	}

	printTable(options)
}

func printTable(options []rainbow.Option) {
	green := color.FgGreen.Render
	red := color.FgRed.Render

	t := newTable(fmt.Sprint("\t\t 29-Oct-21 CeDeFi options: ", len(options)))

	t.AppendHeader(table.Row{"Provider", "Asset", "Type", "Bid size", "Bid Price", "Strike", "Ask Price", "Ask size", "Instrument"})

	for _, option := range options {
		t.AppendRows([]table.Row{{
			provider(option.Provider), option.Asset, option.Type,
			option.Bid[0].Quantity, green(option.Bid[0].Price), option.Strike,
			red(option.Ask[0].Price), option.Ask[0].Quantity, option.Name,
		}})
	}

	t.SortBy([]table.SortBy{
		{Name: "Strike", Mode: table.AscNumeric},
		{Name: "Type", Mode: table.Dsc},
	})

	t.Render()
}

func provider(p string) string {
	magenta := color.FgMagenta.Render
	green := color.FgGreen.Render
	blue := color.FgCyan.Render

	switch p {
	case "Opyn":
		return blue(p)
	case "Deribit":
		return green(p)
	case "PsyOptions":
		return magenta(p)
	default:
		return p
	}
}

func newTable(title string) table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.SetTitle(title)

	return t
}