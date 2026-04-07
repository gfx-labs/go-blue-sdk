package morphosdk

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/holiman/uint256"
)

// Eip5267Domain represents EIP-5267 domain information
type Eip5267Domain struct {
	Fields            hexutil.Bytes  `json:"fields"`
	Name              string         `json:"name"`
	Version           string         `json:"version"`
	ChainId           uint256.Int    `json:"chainId"`
	VerifyingContract common.Address `json:"verifyingContract"`
	Salt              common.Hash    `json:"salt"`
	Extensions        []uint256.Int  `json:"extensions"`
}

// Token represents an ERC20 token
type Token struct {
	Address       common.Address `json:"address"`
	Name          *string        `json:"name,omitempty"`
	Symbol        *string        `json:"symbol,omitempty"`
	Decimals      int            `json:"decimals"`
	Eip5267Domain *Eip5267Domain `json:"eip5267Domain,omitempty"`
	Price         *uint256.Int   `json:"price,omitempty"`
}
