package morphoblue

import "github.com/holiman/uint256"

// WadDivToZero returns (x * WAD) / y, using signed arithmetic
func WadDivToZero(z *uint256.Int, x *uint256.Int, y *uint256.Int) *uint256.Int {
	z = z.Mul(x, WAD)
	z = z.SDiv(z, y)
	return z
}

// WadMulToZero computes x * y / WAD, using signed arithmetic
func WadMulToZero(z *uint256.Int, x *uint256.Int, y *uint256.Int) *uint256.Int {
	z.Mul(x, y)
	z.SDiv(z, WAD)
	return z
}
