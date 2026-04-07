package morphosdk

import (
	"github.com/ethereum/go-ethereum/common"
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
	Address       common.Address `json:"address"`
	Asset         common.Address `json:"asset"`
	Name          string         `json:"name"`
	Symbol        string         `json:"symbol"`
	Decimals      uint8          `json:"decimals"`
	VirtualShares uint256.Int    `json:"virtualShares"`

	TotalAssets uint256.Int `json:"totalAssets"`
	TotalSupply uint256.Int `json:"totalSupply"`

	MaxRate    uint256.Int `json:"maxRate"`
	LastUpdate uint256.Int `json:"lastUpdate"`

	// Adapters
	Adapters         []common.Address `json:"adapters"`
	LiquidityAdapter common.Address   `json:"liquidityAdapter"`
	LiquidityData    []byte           `json:"liquidityData,omitempty"`

	// Allocations (caps per id returned by the liquidity adapter)
	LiquidityAllocations []VaultV2Allocation `json:"liquidityAllocations,omitempty"`

	// Fees
	PerformanceFee          uint256.Int    `json:"performanceFee"`
	ManagementFee           uint256.Int    `json:"managementFee"`
	PerformanceFeeRecipient common.Address `json:"performanceFeeRecipient"`
	ManagementFeeRecipient  common.Address `json:"managementFeeRecipient"`

	// Roles
	Owner   common.Address `json:"owner"`
	Curator common.Address `json:"curator"`
}
