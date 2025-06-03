package morphoblue

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// MarketParamsLib provides functions for working with MarketParams
// This is a Go implementation of the Solidity MarketParamsLib from Morpho Blue

const (
	// MARKET_PARAMS_BYTES_LENGTH represents the total byte length of market parameters (5 * 32 bytes)
	MARKET_PARAMS_BYTES_LENGTH = 5 * 32
)

// ComputeMarketId computes the unique identifier for a market based on its parameters
// This matches the Solidity implementation which uses keccak256 hash of the packed struct
func ComputeMarketId(params MarketParams) common.Hash {
	// Create a byte array to hold all parameters in the correct order
	// The Solidity struct packing order is:
	// 1. loanToken (address = 20 bytes, padded to 32)
	// 2. collateralToken (address = 20 bytes, padded to 32)
	// 3. oracle (address = 20 bytes, padded to 32)
	// 4. irm (address = 20 bytes, padded to 32)
	// 5. lltv (uint256 = 32 bytes)

	data := make([]byte, MARKET_PARAMS_BYTES_LENGTH)

	// Pack loanToken (left-padded with zeros to 32 bytes)
	copy(data[12:32], params.LoanToken.Bytes())

	// Pack collateralToken (left-padded with zeros to 32 bytes)
	copy(data[44:64], params.CollateralToken.Bytes())

	// Pack oracle (left-padded with zeros to 32 bytes)
	copy(data[76:96], params.Oracle.Bytes())

	// Pack irm (left-padded with zeros to 32 bytes)
	copy(data[108:128], params.Irm.Bytes())

	// Pack lltv (already 32 bytes)
	lltvBytes := params.Lltv.Bytes32()
	copy(data[128:160], lltvBytes[:])

	// Compute keccak256 hash
	return crypto.Keccak256Hash(data)
}

// ValidateMarketParams validates that the market parameters are valid
func ValidateMarketParams(params MarketParams) error {
	// Check for zero addresses
	if params.LoanToken == (common.Address{}) {
		return ErrorZeroAddress
	}
	if params.CollateralToken == (common.Address{}) {
		return ErrorZeroAddress
	}
	if params.Oracle == (common.Address{}) {
		return ErrorZeroAddress
	}
	if params.Irm == (common.Address{}) {
		return ErrorZeroAddress
	}

	// Check that loan token and collateral token are different
	if params.LoanToken == params.CollateralToken {
		return ErrorInconsistentInput
	}

	return nil
}

// MarketParamsEqual checks if two MarketParams are equal
func MarketParamsEqual(a, b MarketParams) bool {
	return a.LoanToken == b.LoanToken &&
		a.CollateralToken == b.CollateralToken &&
		a.Oracle == b.Oracle &&
		a.Irm == b.Irm &&
		a.Lltv.Cmp(&b.Lltv) == 0
}

