package morphoblue

import "github.com/holiman/uint256"

func GetUtilization(totalSupplyAssets, totalBorrowAssets *uint256.Int) (*uint256.Int, error) {
	z := new(uint256.Int)
	if totalSupplyAssets.IsZero() {
		if !totalBorrowAssets.IsZero() {
			return z.Set(&MaxUint256), nil
		} else {
			return z.Clear(), nil
		}
	}
	return MulDiv(z, totalBorrowAssets, WAD, totalSupplyAssets)
}

func GetSupplyRate(borrowRate, utilization, fee *uint256.Int) (*uint256.Int, error) {
	z := new(uint256.Int)
	borrowRateWithoutFee, err := WadMulUp(z, borrowRate, utilization)
	if err != nil {
		return nil, err
	}
	w := new(uint256.Int).Set(WAD)
	w = w.Sub(w, fee)

	ans, err := WadMulUp(z, borrowRateWithoutFee, w)
	if err != nil {
		return nil, err
	}
	return ans, nil
}
