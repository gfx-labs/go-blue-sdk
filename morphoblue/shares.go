package morphoblue

import (
	"github.com/holiman/uint256"
)

var (
	virtualShares = uint256.NewInt(1000000)
	virtualAssets = uint256.NewInt(1)
)

func GetAssetsFromShares(shares, totalAssets, totalShares *uint256.Int, roundUp bool) (*uint256.Int, error) {
	if roundUp {
		return MulDivRoundingUp(new(uint256.Int),
			shares,
			new(uint256.Int).Add(totalAssets, virtualAssets),
			new(uint256.Int).Add(totalShares, virtualShares),
		)
	}
	return MulDiv(new(uint256.Int),
		shares,
		new(uint256.Int).Add(totalAssets, virtualAssets),
		new(uint256.Int).Add(totalShares, virtualShares),
	)
}

func GetSharesFromAssets(assets, totalAssets, totalShares *uint256.Int, roundUp bool) (*uint256.Int, error) {
	if roundUp {
		return MulDivRoundingUp(new(uint256.Int),
			assets,
			new(uint256.Int).Add(totalShares, virtualShares),
			new(uint256.Int).Add(totalAssets, virtualAssets),
		)
	}
	return MulDiv(new(uint256.Int),
		assets,
		new(uint256.Int).Add(totalShares, virtualShares),
		new(uint256.Int).Add(totalAssets, virtualAssets),
	)
}
