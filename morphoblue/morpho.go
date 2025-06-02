package morphoblue

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gfx-labs/go-blue-sdk/gosol"
	"github.com/holiman/uint256"
)

// this is supposed to represent the morpho contract
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
