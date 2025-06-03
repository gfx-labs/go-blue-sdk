package morphosdk

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gfx-labs/go-blue-sdk/morphoblue"
)

// ComputeMarketId computes the unique identifier for a market based on its parameters
// This delegates to the morphoblue package implementation
func ComputeMarketId(params MarketParams) common.Hash {
	// Convert morphosdk.MarketParams to morphoblue.MarketParams
	morphoParams := morphoblue.MarketParams{
		LoanToken:       params.LoanToken,
		CollateralToken: params.CollateralToken,
		Oracle:          params.Oracle,
		Irm:             params.Irm,
		Lltv:            params.Lltv,
	}
	return morphoblue.ComputeMarketId(morphoParams)
}