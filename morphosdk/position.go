package morphosdk

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

// Position represents a user's position in a market
type Position struct {
	User          common.Address `json:"user"`
	MarketId      common.Hash    `json:"marketId"`
	SupplyShares  uint256.Int    `json:"supplyShares"`
	BorrowShares  uint256.Int    `json:"borrowShares"`
	Collateral    uint256.Int    `json:"collateral"`
}

// AccrualPosition represents a position with market accrual information
type AccrualPosition struct {
	User          common.Address `json:"user"`
	SupplyShares  uint256.Int    `json:"supplyShares"`
	BorrowShares  uint256.Int    `json:"borrowShares"`
	Collateral    uint256.Int    `json:"collateral"`
	Market        Market         `json:"market"`
}