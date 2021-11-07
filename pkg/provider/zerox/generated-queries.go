package zerox

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// getOptionsOtokensOToken includes the requested fields of the GraphQL type OToken.
type getOptionsOtokensOToken struct {
	Id              string                                      `json:"id"`
	Symbol          string                                      `json:"symbol"`
	Name            string                                      `json:"name"`
	Decimals        int                                         `json:"decimals"`
	StrikeAsset     getOptionsOtokensOTokenStrikeAssetERC20     `json:"strikeAsset"`
	UnderlyingAsset getOptionsOtokensOTokenUnderlyingAssetERC20 `json:"underlyingAsset"`
	CollateralAsset getOptionsOtokensOTokenCollateralAssetERC20 `json:"collateralAsset"`
	StrikePrice     string                                      `json:"strikePrice"`
	IsPut           bool                                        `json:"isPut"`
	ExpiryTimestamp string                                      `json:"expiryTimestamp"`
}

// getOptionsOtokensOTokenCollateralAssetERC20 includes the requested fields of the GraphQL type ERC20.
type getOptionsOtokensOTokenCollateralAssetERC20 struct {
	Id       string `json:"id"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
}

// getOptionsOtokensOTokenStrikeAssetERC20 includes the requested fields of the GraphQL type ERC20.
type getOptionsOtokensOTokenStrikeAssetERC20 struct {
	Id       string `json:"id"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
}

// getOptionsOtokensOTokenUnderlyingAssetERC20 includes the requested fields of the GraphQL type ERC20.
type getOptionsOtokensOTokenUnderlyingAssetERC20 struct {
	Id       string `json:"id"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
}

// getOptionsResponse is returned by getOptions on success.
type getOptionsResponse struct {
	Otokens []getOptionsOtokensOToken `json:"otokens"`
}

func getOptions(
	ctx context.Context,
	client graphql.Client,
) (*getOptionsResponse, error) {
	var err error

	var retval getOptionsResponse
	err = client.MakeRequest(
		ctx,
		"getOptions",
		`
query getOptions {
	otokens {
		id
		symbol
		name
		decimals
		strikeAsset {
			id
			symbol
			decimals
		}
		underlyingAsset {
			id
			symbol
			decimals
		}
		collateralAsset {
			id
			symbol
			decimals
		}
		strikePrice
		isPut
		expiryTimestamp
	}
}
`,
		&retval,
		nil,
	)
	return &retval, err
}