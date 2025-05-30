package morphosdk

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

// User represents a Morpho Blue user
type User struct {
	Address              common.Address `json:"address"`
	IsBundlerAuthorized  bool           `json:"isBundlerAuthorized"`
	MorphoNonce          uint256.Int    `json:"morphoNonce"`
}