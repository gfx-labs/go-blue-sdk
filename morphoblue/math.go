package morphoblue

import "github.com/holiman/uint256"

// WadDivToZero returns (x * WAD) / y, using signed arithmetic
func WadDivToZero(z *uint256.Int, x *uint256.Int, y *uint256.Int) *uint256.Int {
	z = z.Mul(x, WAD)
	z = z.SDiv(z, y)
	return z
}

func WadMulUp(z *uint256.Int, x *uint256.Int, y *uint256.Int) (*uint256.Int, error) {
	_, err := MulDivRoundingUp(z, z, x, WAD)
	return z, err
}

// WadMulToZero computes x * y / WAD, using signed arithmetic
func WadMulToZero(z *uint256.Int, x *uint256.Int, y *uint256.Int) *uint256.Int {
	z.Mul(x, y)
	z.SDiv(z, WAD)
	return z
}

func MulDiv(z *uint256.Int, a, b, denominator *uint256.Int) (*uint256.Int, error) {
	if denominator.IsZero() {
		return nil, ErrorDivideByZero
	}
	z.MulDivOverflow(a, b, denominator)
	return z, nil
}

func MulDivRoundingUp(z *uint256.Int, a, b, denominator *uint256.Int) (*uint256.Int, error) {
	if denominator.IsZero() {
		return nil, ErrorDivideByZero
	}
	if !z.MulMod(a, b, denominator).IsZero() {
		_, err := MulDiv(z, a, b, denominator)
		if err != nil {
			return nil, err
		}
		_, over := z.AddOverflow(z, One)
		if over {
			return nil, ErrorUint256Overflow
		}
	}
	_, err := MulDiv(z, a, b, denominator)
	if err != nil {
		return nil, err
	}
	return z, nil
}
