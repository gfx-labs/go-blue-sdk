package morphosdk

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

// Erc20AllowanceRecipient represents the recipient of an ERC20 allowance
type Erc20AllowanceRecipient string

const (
	Erc20AllowanceRecipientMorpho  Erc20AllowanceRecipient = "morpho"
	Erc20AllowanceRecipientBundler Erc20AllowanceRecipient = "bundler"
	Erc20AllowanceRecipientPermit2 Erc20AllowanceRecipient = "permit2"
)

// Permit2Allowance represents a Permit2 allowance
type Permit2Allowance struct {
	Amount     uint256.Int `json:"amount"`
	Expiration uint256.Int `json:"expiration"`
	Nonce      uint256.Int `json:"nonce"`
}

// Holding represents a user's token holding and allowances
type Holding struct {
	User                    common.Address                          `json:"user"`
	Token                   common.Address                          `json:"token"`
	CanTransfer             *bool                                   `json:"canTransfer,omitempty"`
	Erc20Allowances         map[Erc20AllowanceRecipient]uint256.Int `json:"erc20Allowances"`
	Permit2BundlerAllowance Permit2Allowance                        `json:"permit2BundlerAllowance"`
	Erc2612Nonce            *uint256.Int                            `json:"erc2612Nonce,omitempty"`
	Balance                 uint256.Int                             `json:"balance"`
}