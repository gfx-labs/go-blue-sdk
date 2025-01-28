package morphoblue

import "github.com/holiman/uint256"

var (
	SECONDS_PER_YEAR = uint256.NewInt(31536000)
	WAD              = uint256.MustFromDecimal("1000000000000000000")
)
