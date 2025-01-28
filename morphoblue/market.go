package morphoblue

import "github.com/holiman/uint256"

func GetUtilization(z *uint256.Int, totalSupplyAssets, totalBorrowAssets *uint256.Int) (*uint256.Int, error) {
	if totalSupplyAssets.IsZero() {
		if !totalBorrowAssets.IsZero() {
			return z.Set(&MaxUint256), nil
		} else {
			return z.Clear(), nil
		}
	}
	return MulDiv(z, totalBorrowAssets, WAD, totalSupplyAssets)
}

func MulDiv(z *uint256.Int, a, b, denominator *uint256.Int) (*uint256.Int, error) {
	if denominator.IsZero() {
		return nil, ErrorDivideByZero
	}
	z.MulDivOverflow(a, b, denominator)
	return z, nil
}
