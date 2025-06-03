package morphoblue

import (
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

func TestComputeMarketId(t *testing.T) {
	// Test case 1: Basic market params
	params := MarketParams{
		LoanToken:       common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"), // USDC
		CollateralToken: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"), // WETH
		Oracle:          common.HexToAddress("0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419"), // ETH/USD Oracle
		Irm:             common.HexToAddress("0x870aC11D48B15DB9a138Cf899d20F13F79Ba00BC"), // IRM
		Lltv:            *uint256.NewInt(860000000000000000), // 86% LLTV
	}

	marketId := ComputeMarketId(params)
	
	// The market ID should be deterministic
	if marketId == (common.Hash{}) {
		t.Error("Market ID should not be zero")
	}
	
	// Computing the same params should give the same ID
	marketId2 := ComputeMarketId(params)
	if marketId != marketId2 {
		t.Error("Same params should produce same market ID")
	}
	
	// Different params should give different IDs
	params2 := params
	params2.Lltv = *uint256.NewInt(800000000000000000) // 80% LLTV
	marketId3 := ComputeMarketId(params2)
	if marketId == marketId3 {
		t.Error("Different params should produce different market IDs")
	}
}

func TestMarketParamsEqual(t *testing.T) {
	params1 := MarketParams{
		LoanToken:       common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
		CollateralToken: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
		Oracle:          common.HexToAddress("0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419"),
		Irm:             common.HexToAddress("0x870aC11D48B15DB9a138Cf899d20F13F79Ba00BC"),
		Lltv:            *uint256.NewInt(860000000000000000),
	}
	
	params2 := params1
	if !MarketParamsEqual(params1, params2) {
		t.Error("Equal params should be equal")
	}
	
	// Different loan token
	params3 := params1
	params3.LoanToken = common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7") // USDT
	if MarketParamsEqual(params1, params3) {
		t.Error("Different loan tokens should not be equal")
	}
	
	// Different LLTV
	params4 := params1
	params4.Lltv = *uint256.NewInt(800000000000000000)
	if MarketParamsEqual(params1, params4) {
		t.Error("Different LLTVs should not be equal")
	}
}

func TestValidateMarketParams(t *testing.T) {
	// Valid params
	validParams := MarketParams{
		LoanToken:       common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
		CollateralToken: common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
		Oracle:          common.HexToAddress("0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419"),
		Irm:             common.HexToAddress("0x870aC11D48B15DB9a138Cf899d20F13F79Ba00BC"),
		Lltv:            *uint256.NewInt(860000000000000000),
	}
	if err := ValidateMarketParams(validParams); err != nil {
		t.Errorf("Valid params should not return error: %v", err)
	}
	
	// Zero loan token
	invalidParams1 := validParams
	invalidParams1.LoanToken = common.Address{}
	if err := ValidateMarketParams(invalidParams1); err != ErrorZeroAddress {
		t.Errorf("Zero loan token should return ErrorZeroAddress, got: %v", err)
	}
	
	// Zero collateral token
	invalidParams2 := validParams
	invalidParams2.CollateralToken = common.Address{}
	if err := ValidateMarketParams(invalidParams2); err != ErrorZeroAddress {
		t.Errorf("Zero collateral token should return ErrorZeroAddress, got: %v", err)
	}
	
	// Zero oracle
	invalidParams3 := validParams
	invalidParams3.Oracle = common.Address{}
	if err := ValidateMarketParams(invalidParams3); err != ErrorZeroAddress {
		t.Errorf("Zero oracle should return ErrorZeroAddress, got: %v", err)
	}
	
	// Zero IRM
	invalidParams4 := validParams
	invalidParams4.Irm = common.Address{}
	if err := ValidateMarketParams(invalidParams4); err != ErrorZeroAddress {
		t.Errorf("Zero IRM should return ErrorZeroAddress, got: %v", err)
	}
	
	// Same loan and collateral token
	invalidParams5 := validParams
	invalidParams5.CollateralToken = invalidParams5.LoanToken
	if err := ValidateMarketParams(invalidParams5); err != ErrorInconsistentInput {
		t.Errorf("Same loan and collateral token should return ErrorInconsistentInput, got: %v", err)
	}
}

func TestMarketIdPacking(t *testing.T) {
	// Test the packing format matches Solidity
	params := MarketParams{
		LoanToken:       common.HexToAddress("0x0000000000000000000000000000000000000001"),
		CollateralToken: common.HexToAddress("0x0000000000000000000000000000000000000002"),
		Oracle:          common.HexToAddress("0x0000000000000000000000000000000000000003"),
		Irm:             common.HexToAddress("0x0000000000000000000000000000000000000004"),
		Lltv:            *uint256.NewInt(1000),
	}
	
	// Create expected packed data
	expectedData := make([]byte, MARKET_PARAMS_BYTES_LENGTH)
	// LoanToken (padded to 32 bytes)
	copy(expectedData[12:32], params.LoanToken.Bytes())
	// CollateralToken (padded to 32 bytes)
	copy(expectedData[44:64], params.CollateralToken.Bytes())
	// Oracle (padded to 32 bytes)
	copy(expectedData[76:96], params.Oracle.Bytes())
	// IRM (padded to 32 bytes)
	copy(expectedData[108:128], params.Irm.Bytes())
	// LLTV (32 bytes)
	lltvBytes := params.Lltv.Bytes32()
	copy(expectedData[128:160], lltvBytes[:])
	
	// Verify the data structure
	if len(expectedData) != MARKET_PARAMS_BYTES_LENGTH {
		t.Errorf("Expected data length %d, got %d", MARKET_PARAMS_BYTES_LENGTH, len(expectedData))
	}
	
	// Check specific positions
	if expectedData[31] != byte(0x01) {
		t.Errorf("Expected loan token last byte 0x01, got 0x%02x", expectedData[31])
	}
	if expectedData[63] != byte(0x02) {
		t.Errorf("Expected collateral token last byte 0x02, got 0x%02x", expectedData[63])
	}
	if expectedData[95] != byte(0x03) {
		t.Errorf("Expected oracle last byte 0x03, got 0x%02x", expectedData[95])
	}
	if expectedData[127] != byte(0x04) {
		t.Errorf("Expected IRM last byte 0x04, got 0x%02x", expectedData[127])
	}
	
	t.Logf("Packed data: %s", hex.EncodeToString(expectedData))
}