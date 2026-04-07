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

	AdaptiveCurveIrm common.Address               `json:"adaptiveCurveIrm"`
	MarketIds        []common.Hash                `json:"marketIds"`
	SupplyShares     map[common.Hash]*uint256.Int `json:"supplyShares"`
	// Go-only enrichment fields; not present in TS IVaultV2MorphoMarketV1AdapterV2
	Morpho       common.Address               `json:"morpho,omitempty"`
	SupplyAssets map[common.Hash]*uint256.Int `json:"supplyAssets,omitempty"`
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

// VaultV2AdapterEntry is a tagged-union wrapper for V2 adapter subtypes.
// Exactly one pointer field is non-nil at any time.
type VaultV2AdapterEntry struct {
	MorphoMarketV1AdapterV2 *VaultV2MorphoMarketV1AdapterV2 `json:"-"`
	MorphoMarketV1Adapter   *VaultV2MorphoMarketV1Adapter   `json:"-"`
	MorphoVaultV1Adapter    *VaultV2MorphoVaultV1Adapter    `json:"-"`
	Unknown                 *VaultV2Adapter                 `json:"-"`
}

func (e VaultV2AdapterEntry) MarshalJSON() ([]byte, error) {
	switch {
	case e.MorphoMarketV1AdapterV2 != nil:
		return json.Marshal(e.MorphoMarketV1AdapterV2)
	case e.MorphoMarketV1Adapter != nil:
		return json.Marshal(e.MorphoMarketV1Adapter)
	case e.MorphoVaultV1Adapter != nil:
		return json.Marshal(e.MorphoVaultV1Adapter)
	case e.Unknown != nil:
		return json.Marshal(e.Unknown)
	default:
		return []byte("null"), nil
	}
}

func (e *VaultV2AdapterEntry) UnmarshalJSON(data []byte) error {
	*e = VaultV2AdapterEntry{} // reset all fields to enforce single-variant invariant
	var peek struct {
		Type VaultV2AdapterType `json:"type"`
	}
	if err := json.Unmarshal(data, &peek); err != nil {
		return fmt.Errorf("vault v2 adapter: %w", err)
	}
	switch peek.Type {
	case VaultV2AdapterTypeMorphoMarketV1AdapterV2:
		e.MorphoMarketV1AdapterV2 = new(VaultV2MorphoMarketV1AdapterV2)
		return json.Unmarshal(data, e.MorphoMarketV1AdapterV2)
	case VaultV2AdapterTypeMorphoMarketV1Adapter:
		e.MorphoMarketV1Adapter = new(VaultV2MorphoMarketV1Adapter)
		return json.Unmarshal(data, e.MorphoMarketV1Adapter)
	case VaultV2AdapterTypeMorphoVaultV1Adapter:
		e.MorphoVaultV1Adapter = new(VaultV2MorphoVaultV1Adapter)
		return json.Unmarshal(data, e.MorphoVaultV1Adapter)
	default:
		e.Unknown = new(VaultV2Adapter)
		return json.Unmarshal(data, e.Unknown)
	}
}

// Base returns the shared VaultV2Adapter fields regardless of which variant is set.
func (e *VaultV2AdapterEntry) Base() *VaultV2Adapter {
	switch {
	case e.MorphoMarketV1AdapterV2 != nil:
		return &e.MorphoMarketV1AdapterV2.VaultV2Adapter
	case e.MorphoMarketV1Adapter != nil:
		return &e.MorphoMarketV1Adapter.VaultV2Adapter
	case e.MorphoVaultV1Adapter != nil:
		return &e.MorphoVaultV1Adapter.VaultV2Adapter
	case e.Unknown != nil:
		return e.Unknown
	default:
		return nil
	}
}
