package morphosdk

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

// PendingAddress represents a pending address change
type PendingAddress struct {
	Value   common.Address `json:"value"`
	ValidAt uint256.Int    `json:"validAt"`
}

// PendingTimelock represents a pending timelock change
type PendingTimelock struct {
	Value   uint256.Int `json:"value"`
	ValidAt uint256.Int `json:"validAt"`
}

// VaultPublicAllocatorConfig represents the public allocator configuration
type VaultPublicAllocatorConfig struct {
	Admin      common.Address `json:"admin"`
	Fee        uint256.Int    `json:"fee"`
	AccruedFee uint256.Int    `json:"accruedFee"`
}

// VaultMarketPublicAllocatorConfig represents the public allocator configuration for a specific market
type VaultMarketPublicAllocatorConfig struct {
	Vault    common.Address `json:"vault"`
	MarketId common.Hash    `json:"marketId"`
	MaxIn    uint256.Int    `json:"maxIn"`
	MaxOut   uint256.Int    `json:"maxOut"`
}

// VaultMarketConfig represents the configuration for a market in a vault
type VaultMarketConfig struct {
	Vault                 common.Address                     `json:"vault"`
	MarketId              common.Hash                        `json:"marketId"`
	Cap                   uint256.Int                        `json:"cap"`
	PendingCap            PendingTimelock                    `json:"pendingCap"`
	RemovableAt           uint256.Int                        `json:"removableAt"`
	Enabled               bool                               `json:"enabled"`
	PublicAllocatorConfig *VaultMarketPublicAllocatorConfig `json:"publicAllocatorConfig,omitempty"`
}

// VaultToken represents the ERC20 token properties of a vault
type VaultToken struct {
	Address     common.Address `json:"address"`
	Name        string         `json:"name"`
	Symbol      string         `json:"symbol"`
	Decimals    int            `json:"decimals"`
	TotalSupply uint256.Int    `json:"totalSupply"`
	TotalAssets uint256.Int    `json:"totalAssets"`
}

// Vault represents a Morpho vault
type Vault struct {
	Name                  string                      `json:"name"`
	Symbol                string                      `json:"symbol"`
	Owner                 common.Address              `json:"owner"`
	Curator               common.Address              `json:"curator"`
	Guardian              common.Address              `json:"guardian"`
	SkimRecipient         common.Address              `json:"skimRecipient"`
	FeeRecipient          common.Address              `json:"feeRecipient"`
	Timelock              uint256.Int                 `json:"timelock"`
	Fee                   uint256.Int                 `json:"fee"`
	PendingOwner          common.Address              `json:"pendingOwner"`
	PendingGuardian       PendingAddress              `json:"pendingGuardian"`
	PendingTimelock       PendingTimelock             `json:"pendingTimelock"`
	SupplyQueue           []common.Hash               `json:"supplyQueue"`
	WithdrawQueue         []common.Hash               `json:"withdrawQueue"`
	LastTotalAssets       uint256.Int                 `json:"lastTotalAssets"`
	LostAssets            *uint256.Int                `json:"lostAssets,omitempty"`
	PublicAllocatorConfig *VaultPublicAllocatorConfig `json:"publicAllocatorConfig,omitempty"`
}
