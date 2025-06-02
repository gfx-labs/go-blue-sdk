package morphoblue

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

type Position struct {
	SupplyShares uint256.Int
	BorrowShares uint256.Int
	Collateral   uint256.Int
}

type Market struct {
	TotalSupplyAssets uint256.Int
	TotalSupplyShares uint256.Int
	TotalBorrowAssets uint256.Int
	TotalBorrowShares uint256.Int
	LastUpdate        uint256.Int
	Fee               uint256.Int
}

type MarketParams struct {
	LoanToken       common.Address
	CollateralToken common.Address
	Oracle          common.Address
	Irm             common.Address
	Lltv            uint256.Int
}
