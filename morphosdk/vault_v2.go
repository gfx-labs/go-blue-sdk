package morphosdk

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/holiman/uint256"
)

// VaultV2Allocation represents the cap and allocation for an id in a V2 vault.
type VaultV2Allocation struct {
	Id          common.Hash `json:"id"`
	AbsoluteCap uint256.Int `json:"absoluteCap"`
	RelativeCap uint256.Int `json:"relativeCap"`
	Allocation  uint256.Int `json:"allocation"`
}

// VaultV2 represents a Morpho Vault V2 with adapter-based allocation.
type VaultV2 struct {
	Token

	// IVaultV2 fields
	Asset         common.Address `json:"asset"`
	VirtualShares uint256.Int    `json:"virtualShares"`

	// TotalAssets is the total assets including virtually accrued interest.
	TotalAssets uint256.Int `json:"totalAssets"`
	// RawTotalAssets is the on-chain storage slot value excluding virtually accrued
	// interest, used for interest accrual and fee calculations. The JSON key
	// "_totalAssets" (with leading underscore) matches the TS SDK's IVaultV2
	// interface and the Solidity storage variable name.
	RawTotalAssets uint256.Int `json:"_totalAssets"`
	TotalSupply    uint256.Int `json:"totalSupply"`

	MaxRate    uint256.Int `json:"maxRate"`
	LastUpdate uint256.Int `json:"lastUpdate"`

	// Adapters
	Adapters         []common.Address `json:"adapters"`
	LiquidityAdapter common.Address   `json:"liquidityAdapter"`
	LiquidityData    hexutil.Bytes    `json:"liquidityData,omitempty"`

	// Allocations (caps per id returned by the liquidity adapter)
	LiquidityAllocations []VaultV2Allocation `json:"liquidityAllocations,omitempty"`

	// Fees
	PerformanceFee          uint256.Int    `json:"performanceFee"`
	ManagementFee           uint256.Int    `json:"managementFee"`
	PerformanceFeeRecipient common.Address `json:"performanceFeeRecipient"`
	ManagementFeeRecipient  common.Address `json:"managementFeeRecipient"`

	// Roles (Go-only enrichment fields; not present in TS IVaultV2)
	Owner   common.Address `json:"owner,omitempty"`
	Curator common.Address `json:"curator,omitempty"`
}
