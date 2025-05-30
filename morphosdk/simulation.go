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
	Address            common.Address `json:"address"`
	Vault              common.Address `json:"vault"`
	IsAuthorized       bool           `json:"isAuthorized"`
	AllowedShares      uint256.Int    `json:"allowedShares"`
	AllowedAssets      uint256.Int    `json:"allowedAssets"`
	Shares             uint256.Int    `json:"shares"`
}

// InputSimulationState represents the input state for simulations
type InputSimulationState struct {
	ChainId             int                                                         `json:"chainId"`
	Block               MinimalBlock                                                `json:"block"`
	Global              *GlobalConfig                                               `json:"global,omitempty"`
	Markets             map[common.Hash]*Market                                     `json:"markets,omitempty"`
	Users               map[common.Address]*User                                    `json:"users,omitempty"`
	Tokens              map[common.Address]*Token                                   `json:"tokens,omitempty"`
	Vaults              map[common.Address]*Vault                                   `json:"vaults,omitempty"`
	Positions           map[common.Address]map[common.Hash]*Position               `json:"positions,omitempty"`
	Holdings            map[common.Address]map[common.Address]*Holding             `json:"holdings,omitempty"`
	VaultMarketConfigs  map[common.Address]map[common.Hash]*VaultMarketConfig      `json:"vaultMarketConfigs,omitempty"`
	VaultUsers          map[common.Address]map[common.Address]*VaultUser           `json:"vaultUsers,omitempty"`
}