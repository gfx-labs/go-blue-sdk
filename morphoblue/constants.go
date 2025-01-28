package morphoblue

import (
	"errors"
	"math"

	"github.com/holiman/uint256"
)

var (
	SECONDS_PER_YEAR = uint256.NewInt(31536000)
	WAD              = uint256.MustFromDecimal("1000000000000000000")
	MaxUint256       = uint256.Int{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64}
	MaxUint160       = uint256.Int{math.MaxUint64, math.MaxUint64, math.MaxUint32, 0}
	MaxUint128       = uint256.Int{math.MaxUint64, math.MaxUint64, 0, 0}
	One              = uint256.NewInt(1)
)

var (
	ErrorDivideByZero    = errors.New("divide by zero")
	ErrorUint256Overflow = errors.New("uint256 overflow")
)
