package morphosdk

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

// VaultV2AdapterType identifies the kind of V2 adapter.
type VaultV2AdapterType string

const (
	VaultV2AdapterTypeMorphoMarketV1AdapterV2 VaultV2AdapterType = "MorphoMarketV1AdapterV2"
	VaultV2AdapterTypeMorphoVaultV1Adapter    VaultV2AdapterType = "MorphoVaultV1Adapter"
	VaultV2AdapterTypeUnknown                 VaultV2AdapterType = "Unknown"
)

// VaultV2Adapter represents a generic V2 vault adapter.
type VaultV2Adapter struct {
	Type        VaultV2AdapterType `json:"type"`
	Address     common.Address     `json:"address"`
	ParentVault common.Address     `json:"parentVault"`
	AdapterId   common.Hash        `json:"adapterId"`
	RealAssets  uint256.Int        `json:"realAssets"`
}

// VaultV2MorphoMarketV1AdapterV2 is a V2 adapter that allocates to Morpho Blue markets.
type VaultV2MorphoMarketV1AdapterV2 struct {
	VaultV2Adapter

	Morpho           common.Address              `json:"morpho"`
	AdaptiveCurveIrm common.Address              `json:"adaptiveCurveIrm"`
	MarketIds        []common.Hash               `json:"marketIds"`
	SupplyShares     map[common.Hash]uint256.Int `json:"supplyShares"`
	SupplyAssets     map[common.Hash]uint256.Int `json:"supplyAssets"`
}
