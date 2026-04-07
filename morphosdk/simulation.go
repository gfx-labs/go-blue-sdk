package morphosdk

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

// MinimalBlock represents basic block information
type MinimalBlock struct {
	Number    uint256.Int `json:"number"`
	Timestamp uint256.Int `json:"timestamp"`
}

// GlobalConfig represents global configuration
type GlobalConfig struct {
	FeeRecipient *common.Address `json:"feeRecipient,omitempty"`
}

// VaultUser represents a user's interaction with a vault
type VaultUser struct {
	User        common.Address `json:"user"`
	Vault       common.Address `json:"vault"`
	IsAllocator bool           `json:"isAllocator"`
	Allowance   uint256.Int    `json:"allowance"`
}

// InputSimulationState represents the input state for simulations
type InputSimulationState struct {
	ChainId            int                                                   `json:"chainId"`
	Block              MinimalBlock                                          `json:"block"`
	Global             *GlobalConfig                                         `json:"global,omitempty"`
	Markets            map[common.Hash]*Market                               `json:"markets,omitempty"`
	Users              map[common.Address]*User                              `json:"users,omitempty"`
	Tokens             map[common.Address]*Token                             `json:"tokens,omitempty"`
	Vaults             map[common.Address]*Vault                             `json:"vaults,omitempty"`
	Positions          map[common.Address]map[common.Hash]*Position          `json:"positions,omitempty"`
	Holdings           map[common.Address]map[common.Address]*Holding        `json:"holdings,omitempty"`
	VaultMarketConfigs map[common.Address]map[common.Hash]*VaultMarketConfig `json:"vaultMarketConfigs,omitempty"`
	VaultUsers         map[common.Address]map[common.Address]*VaultUser      `json:"vaultUsers,omitempty"`

	// V2 vault state
	VaultV2s        map[common.Address]*VaultV2             `json:"vaultV2s,omitempty"`
	VaultV2Adapters map[common.Address]*VaultV2AdapterEntry `json:"vaultV2Adapters,omitempty"`
}
