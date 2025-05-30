package morphosdk

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

// MarketParams represents the parameters defining a market
type MarketParams struct {
	LoanToken       common.Address `json:"loanToken"`
	CollateralToken common.Address `json:"collateralToken"`
	Oracle          common.Address `json:"oracle"`
	Irm             common.Address `json:"irm"`
	Lltv            uint256.Int    `json:"lltv"`
}

// Market represents a Morpho Blue market
type Market struct {
	Params            MarketParams `json:"params"`
	TotalSupplyAssets uint256.Int  `json:"totalSupplyAssets"`
	TotalBorrowAssets uint256.Int  `json:"totalBorrowAssets"`
	TotalSupplyShares uint256.Int  `json:"totalSupplyShares"`
	TotalBorrowShares uint256.Int  `json:"totalBorrowShares"`
	LastUpdate        uint256.Int  `json:"lastUpdate"`
	Fee               uint256.Int  `json:"fee"`
	Price             *uint256.Int `json:"price,omitempty"`
	RateAtTarget      *uint256.Int `json:"rateAtTarget,omitempty"`
}

// CapacityLimit represents a capacity limit with its reason
type CapacityLimit struct {
	Value   uint256.Int         `json:"value"`
	Limiter CapacityLimitReason `json:"limiter"`
}

// MaxBorrowOptions represents options for maximum borrow calculations
type MaxBorrowOptions struct {
	MaxLtv *uint256.Int `json:"maxLtv,omitempty"`
}

// MaxWithdrawCollateralOptions represents options for maximum collateral withdrawal
type MaxWithdrawCollateralOptions struct {
	MaxLtv *uint256.Int `json:"maxLtv,omitempty"`
}

// MaxPositionCapacities represents the maximum capacities for different position operations
type MaxPositionCapacities struct {
	Supply             CapacityLimit  `json:"supply"`
	Withdraw           CapacityLimit  `json:"withdraw"`
	Borrow             *CapacityLimit `json:"borrow,omitempty"`
	Repay              CapacityLimit  `json:"repay"`
	SupplyCollateral   CapacityLimit  `json:"supplyCollateral"`
	WithdrawCollateral *CapacityLimit `json:"withdrawCollateral,omitempty"`
}