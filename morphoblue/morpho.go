package morphoblue

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gfx-labs/go-blue-sdk/gosol"
	"github.com/holiman/uint256"
)

// this is supposed to represent the morpho contract
// reference implementation:
// https://github.com/morpho-org/morpho-blue/blob/main/src/Morpho.sol
type Morpho struct {
	Owner        common.Address
	FeeRecipient common.Address

	Position gosol.Mapping[common.Hash, gosol.Mapping[common.Address, Position]]
	Market   gosol.Mapping[common.Hash, Market]

	IsIrmEnabled  gosol.Mapping[common.Address, bool]
	IsLltvEnabled gosol.Mapping[common.Address, bool]

	// NOTE: we don't care about this for now.
	// IsAuthorized gosol.Mapping[common.Address, bool]

	Nonce gosol.Mapping[common.Address, uint256.Int]

	IdToMarketParams gosol.Mapping[common.Hash, MarketParams]
}

// NewMorpho creates a new Morpho instance with the given owner and fee recipient
func NewMorpho(owner, feeRecipient common.Address) *Morpho {
	return &Morpho{
		Owner:            owner,
		FeeRecipient:     feeRecipient,
		Position:         gosol.NewMapMapping[common.Hash, gosol.Mapping[common.Address, Position]](),
		Market:           gosol.NewMapMapping[common.Hash, Market](),
		IsIrmEnabled:     gosol.NewMapMapping[common.Address, bool](),
		IsLltvEnabled:    gosol.NewMapMapping[common.Address, bool](),
		Nonce:            gosol.NewMapMapping[common.Address, uint256.Int](),
		IdToMarketParams: gosol.NewMapMapping[common.Hash, MarketParams](),
	}
}
