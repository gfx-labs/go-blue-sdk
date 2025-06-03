package morphoblue

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
)

// Mock IRM for testing
type mockIRM struct {
	rate *uint256.Int
}

func (m *mockIRM) BorrowRate(marketParams MarketParams, market Market) (*uint256.Int, error) {
	return m.rate, nil
}

func TestMorphoBasicOperations(t *testing.T) {
	// Create a new Morpho instance
	owner := common.HexToAddress("0x1234567890123456789012345678901234567890")
	feeRecipient := common.HexToAddress("0x0987654321098765432109876543210987654321")
	morpho := NewMorpho(owner, feeRecipient)

	// Test owner setup
	require.Equal(t, owner, morpho.Owner)
	require.Equal(t, feeRecipient, morpho.FeeRecipient)

	// Enable IRM and LLTV
	irmAddr := common.HexToAddress("0x1111111111111111111111111111111111111111")
	lltv := uint256.MustFromDecimal("800000000000000000") // 80%

	err := morpho.EnableIrm(owner, irmAddr)
	require.NoError(t, err)

	err = morpho.EnableLltv(owner, lltv)
	require.NoError(t, err)

	// Set up oracle address
	oracleAddr := common.HexToAddress("0x2222222222222222222222222222222222222222")

	// Create market
	marketParams := MarketParams{
		LoanToken:       common.HexToAddress("0x3333333333333333333333333333333333333333"),
		CollateralToken: common.HexToAddress("0x4444444444444444444444444444444444444444"),
		Oracle:          oracleAddr,
		Irm:             irmAddr,
		Lltv:            *lltv,
	}

	err = morpho.CreateMarket(owner, marketParams)
	require.NoError(t, err)

	// Test market ID computation
	marketId := morpho.GetMarketId(marketParams)
	require.NotEqual(t, common.Hash{}, marketId)

	// Verify market was created
	market, _ := morpho.Market.Get(marketId)
	require.False(t, market.LastUpdate.IsZero())
}

func TestMorphoSupplyWithdraw(t *testing.T) {
	// Setup
	owner := common.HexToAddress("0x1234567890123456789012345678901234567890")
	morpho := NewMorpho(owner, owner)
	
	irmAddr := common.HexToAddress("0x1111111111111111111111111111111111111111")
	morpho.EnableIrm(owner, irmAddr)
	
	lltv := uint256.MustFromDecimal("800000000000000000")
	morpho.EnableLltv(owner, lltv)
	
	marketParams := MarketParams{
		LoanToken:       common.HexToAddress("0x3333333333333333333333333333333333333333"),
		CollateralToken: common.HexToAddress("0x4444444444444444444444444444444444444444"),
		Oracle:          common.HexToAddress("0x2222222222222222222222222222222222222222"),
		Irm:             irmAddr,
		Lltv:            *lltv,
	}
	
	morpho.CreateMarket(owner, marketParams)
	
	// Test supply
	supplier := common.HexToAddress("0x5555555555555555555555555555555555555555")
	supplyAmount := uint256.MustFromDecimal("1000000000000000000") // 1 token
	borrowRate := uint256.MustFromDecimal("50000000000000000") // 5%
	
	assets, shares, err := morpho.Supply(supplier, uint256.NewInt(0), marketParams, supplyAmount, uint256.NewInt(0), supplier, nil, borrowRate)
	require.NoError(t, err)
	require.NotNil(t, assets)
	require.NotNil(t, shares)
	require.True(t, assets.Cmp(uint256.NewInt(0)) > 0)
	require.True(t, shares.Cmp(uint256.NewInt(0)) > 0)
	
	// Verify position
	marketId := morpho.GetMarketId(marketParams)
	position, _ := morpho.getPosition(marketId, supplier)
	require.True(t, position.SupplyShares.Cmp(uint256.NewInt(0)) > 0)
	
	// Test withdraw
	collateralPrice := WAD // 1:1 price
	withdrawAssets, withdrawShares, err := morpho.Withdraw(supplier, marketParams, assets, uint256.NewInt(0), supplier, supplier, borrowRate, collateralPrice)
	require.NoError(t, err)
	require.Equal(t, assets.String(), withdrawAssets.String())
	require.Equal(t, shares.String(), withdrawShares.String())
	
	// Verify position after withdraw
	position, _ = morpho.getPosition(marketId, supplier)
	require.True(t, position.SupplyShares.IsZero())
}