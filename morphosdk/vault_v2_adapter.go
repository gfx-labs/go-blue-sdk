package morphosdk

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

// VaultV2AdapterType identifies the kind of V2 adapter.
type VaultV2AdapterType string

const (
	VaultV2AdapterTypeMorphoMarketV1AdapterV2 VaultV2AdapterType = "VaultV2MorphoMarketV1AdapterV2"
	VaultV2AdapterTypeMorphoMarketV1Adapter   VaultV2AdapterType = "VaultV2MorphoMarketV1Adapter"
	VaultV2AdapterTypeMorphoVaultV1Adapter    VaultV2AdapterType = "VaultV2MorphoVaultV1Adapter"
	VaultV2AdapterTypeUnknown                 VaultV2AdapterType = "Unknown"
)

// VaultV2Adapter represents the base fields shared by all V2 vault adapters.
type VaultV2Adapter struct {
	Type          VaultV2AdapterType `json:"type"`
	Address       common.Address     `json:"address"`
	ParentVault   common.Address     `json:"parentVault"`
	AdapterId     common.Hash        `json:"adapterId"`
	SkimRecipient common.Address     `json:"skimRecipient"`
	// RealAssets is a point-in-time snapshot of the adapter's total asset value.
	// In the TS SDK this is a method on IAccrualVaultV2Adapter; here it is stored
	// as a snapshot value from the last on-chain query.
	RealAssets uint256.Int `json:"realAssets"`
}

// VaultV2MorphoMarketV1AdapterV2 is the v2 adapter that allocates to Morpho Blue markets.
// It uses marketIds and per-market supplyShares for position tracking.
type VaultV2MorphoMarketV1AdapterV2 struct {
	VaultV2Adapter

	AdaptiveCurveIrm common.Address              `json:"adaptiveCurveIrm"`
	MarketIds        []common.Hash               `json:"marketIds"`
	SupplyShares     map[common.Hash]uint256.Int `json:"supplyShares"`
	// Go-only enrichment fields; not present in TS IVaultV2MorphoMarketV1AdapterV2
	Morpho       common.Address              `json:"morpho,omitempty"`
	SupplyAssets map[common.Hash]uint256.Int `json:"supplyAssets,omitempty"`
}

// VaultV2MorphoMarketV1Adapter is the v1 adapter that allocates to Morpho Blue markets.
// It uses full MarketParams objects rather than just market IDs.
type VaultV2MorphoMarketV1Adapter struct {
	VaultV2Adapter

	MarketParamsList []MarketParams `json:"marketParamsList"`
}

// VaultV2MorphoVaultV1Adapter is an adapter that wraps a V1 MetaMorpho vault as a liquidity source.
type VaultV2MorphoVaultV1Adapter struct {
	VaultV2Adapter

	MorphoVaultV1 common.Address `json:"morphoVaultV1"`
}

// VaultV2AdapterEntry is a polymorphic wrapper that serializes any V2 adapter subtype
// to JSON and deserializes it back using the "type" field as a discriminator.
// Use this in maps/slices where mixed adapter types must roundtrip through JSON.
type VaultV2AdapterEntry struct {
	// Adapter holds the concrete adapter value. It must be one of:
	//   *VaultV2MorphoMarketV1AdapterV2
	//   *VaultV2MorphoMarketV1Adapter
	//   *VaultV2MorphoVaultV1Adapter
	//   *VaultV2Adapter (for unknown types)
	Adapter any
}

func (e VaultV2AdapterEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Adapter)
}

func (e *VaultV2AdapterEntry) UnmarshalJSON(data []byte) error {
	// Peek at the type field to determine the concrete type.
	var peek struct {
		Type VaultV2AdapterType `json:"type"`
	}
	if err := json.Unmarshal(data, &peek); err != nil {
		return fmt.Errorf("vault v2 adapter: %w", err)
	}
	switch peek.Type {
	case VaultV2AdapterTypeMorphoMarketV1AdapterV2:
		var v VaultV2MorphoMarketV1AdapterV2
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		e.Adapter = &v
	case VaultV2AdapterTypeMorphoMarketV1Adapter:
		var v VaultV2MorphoMarketV1Adapter
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		e.Adapter = &v
	case VaultV2AdapterTypeMorphoVaultV1Adapter:
		var v VaultV2MorphoVaultV1Adapter
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		e.Adapter = &v
	default:
		var v VaultV2Adapter
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		e.Adapter = &v
	}
	return nil
}

// Base returns the base VaultV2Adapter fields regardless of the concrete type.
func (e *VaultV2AdapterEntry) Base() *VaultV2Adapter {
	switch v := e.Adapter.(type) {
	case *VaultV2MorphoMarketV1AdapterV2:
		return &v.VaultV2Adapter
	case *VaultV2MorphoMarketV1Adapter:
		return &v.VaultV2Adapter
	case *VaultV2MorphoVaultV1Adapter:
		return &v.VaultV2Adapter
	case *VaultV2Adapter:
		return v
	default:
		return nil
	}
}
