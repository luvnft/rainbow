// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package opyn

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// OptionsOtokensOToken includes the requested fields of the GraphQL type OToken.
type OptionsOtokensOToken struct {
	Id              string                                   `json:"id"`
	Symbol          string                                   `json:"symbol"`
	Name            string                                   `json:"name"`
	Decimals        int                                      `json:"decimals"`
	StrikeAsset     OptionsOtokensOTokenStrikeAssetERC20     `json:"strikeAsset"`
	UnderlyingAsset OptionsOtokensOTokenUnderlyingAssetERC20 `json:"underlyingAsset"`
	CollateralAsset OptionsOtokensOTokenCollateralAssetERC20 `json:"collateralAsset"`
	StrikePrice     string                                   `json:"strikePrice"`
	IsPut           bool                                     `json:"isPut"`
	ExpiryTimestamp string                                   `json:"expiryTimestamp"`
}

// GetId returns OptionsOtokensOToken.Id, and is useful for accessing the field via an interface.
func (v *OptionsOtokensOToken) GetId() string { return v.Id }

// GetSymbol returns OptionsOtokensOToken.Symbol, and is useful for accessing the field via an interface.
func (v *OptionsOtokensOToken) GetSymbol() string { return v.Symbol }

// GetName returns OptionsOtokensOToken.Name, and is useful for accessing the field via an interface.
func (v *OptionsOtokensOToken) GetName() string { return v.Name }

// GetDecimals returns OptionsOtokensOToken.Decimals, and is useful for accessing the field via an interface.
func (v *OptionsOtokensOToken) GetDecimals() int { return v.Decimals }

// GetStrikeAsset returns OptionsOtokensOToken.StrikeAsset, and is useful for accessing the field via an interface.
func (v *OptionsOtokensOToken) GetStrikeAsset() OptionsOtokensOTokenStrikeAssetERC20 {
	return v.StrikeAsset
}

// GetUnderlyingAsset returns OptionsOtokensOToken.UnderlyingAsset, and is useful for accessing the field via an interface.
func (v *OptionsOtokensOToken) GetUnderlyingAsset() OptionsOtokensOTokenUnderlyingAssetERC20 {
	return v.UnderlyingAsset
}

// GetCollateralAsset returns OptionsOtokensOToken.CollateralAsset, and is useful for accessing the field via an interface.
func (v *OptionsOtokensOToken) GetCollateralAsset() OptionsOtokensOTokenCollateralAssetERC20 {
	return v.CollateralAsset
}

// GetStrikePrice returns OptionsOtokensOToken.StrikePrice, and is useful for accessing the field via an interface.
func (v *OptionsOtokensOToken) GetStrikePrice() string { return v.StrikePrice }

// GetIsPut returns OptionsOtokensOToken.IsPut, and is useful for accessing the field via an interface.
func (v *OptionsOtokensOToken) GetIsPut() bool { return v.IsPut }

// GetExpiryTimestamp returns OptionsOtokensOToken.ExpiryTimestamp, and is useful for accessing the field via an interface.
func (v *OptionsOtokensOToken) GetExpiryTimestamp() string { return v.ExpiryTimestamp }

// OptionsOtokensOTokenCollateralAssetERC20 includes the requested fields of the GraphQL type ERC20.
type OptionsOtokensOTokenCollateralAssetERC20 struct {
	Id       string `json:"id"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
}

// GetId returns OptionsOtokensOTokenCollateralAssetERC20.Id, and is useful for accessing the field via an interface.
func (v *OptionsOtokensOTokenCollateralAssetERC20) GetId() string { return v.Id }

// GetSymbol returns OptionsOtokensOTokenCollateralAssetERC20.Symbol, and is useful for accessing the field via an interface.
func (v *OptionsOtokensOTokenCollateralAssetERC20) GetSymbol() string { return v.Symbol }

// GetDecimals returns OptionsOtokensOTokenCollateralAssetERC20.Decimals, and is useful for accessing the field via an interface.
func (v *OptionsOtokensOTokenCollateralAssetERC20) GetDecimals() int { return v.Decimals }

// OptionsOtokensOTokenStrikeAssetERC20 includes the requested fields of the GraphQL type ERC20.
type OptionsOtokensOTokenStrikeAssetERC20 struct {
	Id       string `json:"id"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
}

// GetId returns OptionsOtokensOTokenStrikeAssetERC20.Id, and is useful for accessing the field via an interface.
func (v *OptionsOtokensOTokenStrikeAssetERC20) GetId() string { return v.Id }

// GetSymbol returns OptionsOtokensOTokenStrikeAssetERC20.Symbol, and is useful for accessing the field via an interface.
func (v *OptionsOtokensOTokenStrikeAssetERC20) GetSymbol() string { return v.Symbol }

// GetDecimals returns OptionsOtokensOTokenStrikeAssetERC20.Decimals, and is useful for accessing the field via an interface.
func (v *OptionsOtokensOTokenStrikeAssetERC20) GetDecimals() int { return v.Decimals }

// OptionsOtokensOTokenUnderlyingAssetERC20 includes the requested fields of the GraphQL type ERC20.
type OptionsOtokensOTokenUnderlyingAssetERC20 struct {
	Id       string `json:"id"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
}

// GetId returns OptionsOtokensOTokenUnderlyingAssetERC20.Id, and is useful for accessing the field via an interface.
func (v *OptionsOtokensOTokenUnderlyingAssetERC20) GetId() string { return v.Id }

// GetSymbol returns OptionsOtokensOTokenUnderlyingAssetERC20.Symbol, and is useful for accessing the field via an interface.
func (v *OptionsOtokensOTokenUnderlyingAssetERC20) GetSymbol() string { return v.Symbol }

// GetDecimals returns OptionsOtokensOTokenUnderlyingAssetERC20.Decimals, and is useful for accessing the field via an interface.
func (v *OptionsOtokensOTokenUnderlyingAssetERC20) GetDecimals() int { return v.Decimals }

// OptionsResponse is returned by Options on success.
type OptionsResponse struct {
	Otokens []OptionsOtokensOToken `json:"otokens"`
}

// GetOtokens returns OptionsResponse.Otokens, and is useful for accessing the field via an interface.
func (v *OptionsResponse) GetOtokens() []OptionsOtokensOToken { return v.Otokens }

// __OptionsInput is used internally by genqlient
type __OptionsInput struct {
	Skip  int    `json:"skip"`
	First int    `json:"first"`
	T     string `json:"t"`
}

// GetSkip returns __OptionsInput.Skip, and is useful for accessing the field via an interface.
func (v *__OptionsInput) GetSkip() int { return v.Skip }

// GetFirst returns __OptionsInput.First, and is useful for accessing the field via an interface.
func (v *__OptionsInput) GetFirst() int { return v.First }

// GetT returns __OptionsInput.T, and is useful for accessing the field via an interface.
func (v *__OptionsInput) GetT() string { return v.T }

func Options(
	ctx context.Context,
	client graphql.Client,
	skip int,
	first int,
	t string,
) (*OptionsResponse, error) {
	__input := __OptionsInput{
		Skip:  skip,
		First: first,
		T:     t,
	}
	var err error

	var retval OptionsResponse
	err = client.MakeRequest(
		ctx,
		"Options",
		`
query Options ($skip: Int, $first: Int, $t: BigInt) {
	otokens(skip: $skip, first: $first, orderBy: expiryTimestamp, orderDirection: desc, where: {expiryTimestamp_gt:$t}) {
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
		&__input,
	)
	return &retval, err
}
