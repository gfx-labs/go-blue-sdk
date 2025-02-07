package morphoblue

import (
	"testing"

	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
)

func TestMaxValues(t *testing.T) {
	require.Equal(t, MaxUint256.String(), "115792089237316195423570985008687907853269984665640564039457584007913129639935")
	require.Equal(t, MaxUint160.String(), "1461501637330902918203684832716283019655932542975")
	require.Equal(t, MaxUint128.String(), "340282366920938463463374607431768211455")
}

func TestWadTaylorCompounded(t *testing.T) {
	type testcase struct {
		x        *uint256.Int
		n        *uint256.Int
		expected *uint256.Int
	}
	testcases := []testcase{
		{
			x:        uint256.NewInt(2),
			n:        uint256.NewInt(5),
			expected: uint256.NewInt(10),
		},
		{
			x:        uint256.NewInt(4214),
			n:        uint256.NewInt(244),
			expected: uint256.NewInt(1028216),
		},
	}
	for _, tc := range testcases {
		result := WadTaylorCompounded(new(uint256.Int), tc.x, tc.n)
		require.Equal(t, result.String(), tc.expected.String())
	}
}
