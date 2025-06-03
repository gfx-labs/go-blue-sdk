package morphoblue

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
)

func TestIsHealthy(t *testing.T) {
	// Create a Morpho instance
	owner := common.HexToAddress("0x1234567890123456789012345678901234567890")
	morpho := NewMorpho(owner, owner)
	
	// Create market parameters with 80% LLTV
	marketParams := MarketParams{
		LoanToken:       common.HexToAddress("0x1111111111111111111111111111111111111111"),
		CollateralToken: common.HexToAddress("0x2222222222222222222222222222222222222222"),
		Oracle:          common.HexToAddress("0x3333333333333333333333333333333333333333"),
		Irm:             common.HexToAddress("0x4444444444444444444444444444444444444444"),
		Lltv:            *uint256.MustFromDecimal("800000000000000000"), // 80%
	}
	
	// Enable IRM and LLTV
	morpho.EnableIrm(owner, marketParams.Irm)
	morpho.EnableLltv(owner, &marketParams.Lltv)
	
	// Create market
	err := morpho.CreateMarket(owner, marketParams)
	require.NoError(t, err)
	
	marketId := morpho.GetMarketId(marketParams)
	
	// Set up market state
	market := Market{
		TotalSupplyAssets: *uint256.MustFromDecimal("1000000000000000000000"), // 1000 tokens
		TotalSupplyShares: *uint256.MustFromDecimal("1000000000000000000000"), // 1000 shares
		TotalBorrowAssets: *uint256.MustFromDecimal("500000000000000000000"),  // 500 tokens
		TotalBorrowShares: *uint256.MustFromDecimal("500000000000000000000"),  // 500 shares
		LastUpdate:        *uint256.NewInt(morpho.BlockTimestamp),
		Fee:               *uint256.NewInt(0),
	}
	morpho.Market.Set(marketId, market)

	testCases := []struct {
		name            string
		user            common.Address
		position        Position
		collateralPrice *uint256.Int
		expectedHealthy bool
	}{
		{
			name: "No debt is always healthy",
			user: common.HexToAddress("0x0000000000000000000000000000000000000001"),
			position: Position{
				SupplyShares: *uint256.MustFromDecimal("100000000000000000000"),
				BorrowShares: *uint256.NewInt(0),
				Collateral:   *uint256.MustFromDecimal("100000000000000000000"),
			},
			collateralPrice: WAD, // 1:1 price
			expectedHealthy: true,
		},
		{
			name: "Healthy position at 50% utilization",
			user: common.HexToAddress("0x0000000000000000000000000000000000000002"),
			position: Position{
				SupplyShares: *uint256.NewInt(0),
				BorrowShares: *uint256.MustFromDecimal("100000000000000000000"), // 100 tokens borrowed
				Collateral:   *uint256.MustFromDecimal("250000000000000000000"), // 250 collateral
			},
			collateralPrice: WAD, // 1:1 price
			expectedHealthy: true, // 100 borrowed < 250 * 0.8 = 200 max borrowable
		},
		{
			name: "Unhealthy position at 90% utilization",
			user: common.HexToAddress("0x0000000000000000000000000000000000000003"),
			position: Position{
				SupplyShares: *uint256.NewInt(0),
				BorrowShares: *uint256.MustFromDecimal("100000000000000000000"), // 100 tokens borrowed
				Collateral:   *uint256.MustFromDecimal("111111111111111111111"), // 111.11 collateral
			},
			collateralPrice: WAD, // 1:1 price
			expectedHealthy: false, // 100 borrowed > 111.11 * 0.8 = 88.88 max borrowable
		},
		{
			name: "Healthy with higher collateral price",
			user: common.HexToAddress("0x0000000000000000000000000000000000000004"),
			position: Position{
				SupplyShares: *uint256.NewInt(0),
				BorrowShares: *uint256.MustFromDecimal("100000000000000000000"), // 100 tokens borrowed
				Collateral:   *uint256.MustFromDecimal("50000000000000000000"),  // 50 collateral
			},
			collateralPrice: uint256.MustFromDecimal("3000000000000000000"), // 3:1 price
			expectedHealthy: true, // 100 borrowed < 50 * 3 * 0.8 = 120 max borrowable
		},
		{
			name: "Unhealthy with lower collateral price",
			user: common.HexToAddress("0x0000000000000000000000000000000000000005"),
			position: Position{
				SupplyShares: *uint256.NewInt(0),
				BorrowShares: *uint256.MustFromDecimal("100000000000000000000"), // 100 tokens borrowed
				Collateral:   *uint256.MustFromDecimal("200000000000000000000"), // 200 collateral
			},
			collateralPrice: uint256.MustFromDecimal("500000000000000000"), // 0.5:1 price
			expectedHealthy: false, // 100 borrowed > 200 * 0.5 * 0.8 = 80 max borrowable
		},
		{
			name: "Exactly at liquidation threshold",
			user: common.HexToAddress("0x0000000000000000000000000000000000000006"),
			position: Position{
				SupplyShares: *uint256.NewInt(0),
				BorrowShares: *uint256.MustFromDecimal("80000000000000000000"), // 80 tokens borrowed
				Collateral:   *uint256.MustFromDecimal("100000000000000000000"), // 100 collateral
			},
			collateralPrice: WAD, // 1:1 price
			expectedHealthy: true, // 80 borrowed = 100 * 0.8 = 80 max borrowable (exactly at limit)
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Set up the position for the test user
			morpho.setPosition(marketId, tc.user, tc.position)
			
			// Test the IsHealthy method
			healthy := morpho.IsHealthy(marketParams, marketId, tc.user, tc.collateralPrice)
			require.Equal(t, tc.expectedHealthy, healthy, "Health check mismatch")
		})
	}
}

func TestCheckHealthIntegration(t *testing.T) {
	// Setup Morpho instance
	owner := common.HexToAddress("0x1234567890123456789012345678901234567890")
	morpho := NewMorpho(owner, owner)
	
	// Enable IRM and LLTV
	irmAddr := common.HexToAddress("0x1111111111111111111111111111111111111111")
	err := morpho.EnableIrm(owner, irmAddr)
	require.NoError(t, err)
	
	lltv := uint256.MustFromDecimal("800000000000000000") // 80%
	err = morpho.EnableLltv(owner, lltv)
	require.NoError(t, err)
	
	// Create market
	marketParams := MarketParams{
		LoanToken:       common.HexToAddress("0x3333333333333333333333333333333333333333"),
		CollateralToken: common.HexToAddress("0x4444444444444444444444444444444444444444"),
		Oracle:          common.HexToAddress("0x2222222222222222222222222222222222222222"),
		Irm:             irmAddr,
		Lltv:            *lltv,
	}
	
	err = morpho.CreateMarket(owner, marketParams)
	require.NoError(t, err)
	
	// Supply liquidity
	supplier := common.HexToAddress("0x5555555555555555555555555555555555555555")
	supplyAmount := uint256.MustFromDecimal("1000000000000000000000") // 1000 tokens
	borrowRate := uint256.MustFromDecimal("50000000000000000") // 5%
	
	_, _, err = morpho.Supply(supplier, uint256.NewInt(0), marketParams, supplyAmount, uint256.NewInt(0), supplier, nil, borrowRate)
	require.NoError(t, err)
	
	// Supply collateral
	borrower := common.HexToAddress("0x6666666666666666666666666666666666666666")
	collateralAmount := uint256.MustFromDecimal("100000000000000000000") // 100 collateral
	
	err = morpho.SupplyCollateral(borrower, uint256.NewInt(0), marketParams, collateralAmount, borrower, nil)
	require.NoError(t, err)
	
	// Test borrowing with different collateral prices
	t.Run("Borrow succeeds with high collateral price", func(t *testing.T) {
		borrowAmount := uint256.MustFromDecimal("50000000000000000000") // 50 tokens
		collateralPrice := uint256.MustFromDecimal("2000000000000000000") // 2:1 price
		
		_, _, err := morpho.Borrow(borrower, marketParams, borrowAmount, uint256.NewInt(0), borrower, borrower, borrowRate, collateralPrice)
		require.NoError(t, err)
	})
	
	// Test withdrawing collateral with health check
	t.Run("Withdraw collateral fails when it would make position unhealthy", func(t *testing.T) {
		withdrawAmount := uint256.MustFromDecimal("90000000000000000000") // Try to withdraw 90 collateral
		collateralPrice := WAD // 1:1 price
		
		err := morpho.WithdrawCollateral(borrower, marketParams, withdrawAmount, borrower, borrower, borrowRate, collateralPrice)
		require.Error(t, err)
		require.Equal(t, ErrorInsufficientCollateral, err)
	})
}