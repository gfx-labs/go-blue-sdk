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

func TestMulDivRoundingUp(t *testing.T) {
	type testcase struct {
		name        string
		a           *uint256.Int
		b           *uint256.Int
		denominator *uint256.Int
		expected    *uint256.Int
		expectError bool
	}

	testcases := []testcase{
		// Exact division cases (no remainder)
		{
			name:        "exact division small numbers",
			a:           uint256.NewInt(10),
			b:           uint256.NewInt(20),
			denominator: uint256.NewInt(10),
			expected:    uint256.NewInt(20),
			expectError: false,
		},
		{
			name:        "exact division larger numbers",
			a:           uint256.NewInt(1000),
			b:           uint256.NewInt(500),
			denominator: uint256.NewInt(100),
			expected:    uint256.NewInt(5000),
			expectError: false,
		},
		{
			name:        "exact division with 1 as denominator",
			a:           uint256.NewInt(123),
			b:           uint256.NewInt(456),
			denominator: uint256.NewInt(1),
			expected:    uint256.NewInt(56088),
			expectError: false,
		},
		// Cases with remainder (should round up)
		{
			name:        "small remainder rounds up",
			a:           uint256.NewInt(10),
			b:           uint256.NewInt(3),
			denominator: uint256.NewInt(7),
			expected:    uint256.NewInt(5), // 30/7 = 4.28... rounds up to 5
			expectError: false,
		},
		{
			name:        "larger remainder rounds up",
			a:           uint256.NewInt(100),
			b:           uint256.NewInt(7),
			denominator: uint256.NewInt(13),
			expected:    uint256.NewInt(54), // 700/13 = 53.84... rounds up to 54
			expectError: false,
		},
		{
			name:        "minimal remainder still rounds up",
			a:           uint256.NewInt(10),
			b:           uint256.NewInt(10),
			denominator: uint256.NewInt(99),
			expected:    uint256.NewInt(2), // 100/99 = 1.01... rounds up to 2
			expectError: false,
		},
		// Edge cases
		{
			name:        "zero numerator",
			a:           uint256.NewInt(0),
			b:           uint256.NewInt(100),
			denominator: uint256.NewInt(10),
			expected:    uint256.NewInt(0),
			expectError: false,
		},
		{
			name:        "one factor is zero",
			a:           uint256.NewInt(100),
			b:           uint256.NewInt(0),
			denominator: uint256.NewInt(10),
			expected:    uint256.NewInt(0),
			expectError: false,
		},
		{
			name:        "divide by zero error",
			a:           uint256.NewInt(10),
			b:           uint256.NewInt(20),
			denominator: uint256.NewInt(0),
			expected:    nil,
			expectError: true,
		},
		// Large number tests
		{
			name:        "large numbers exact division",
			a:           uint256.MustFromDecimal("1000000000000000000"),
			b:           uint256.MustFromDecimal("2000000000000000000"),
			denominator: uint256.MustFromDecimal("1000000000000000000"),
			expected:    uint256.MustFromDecimal("2000000000000000000"),
			expectError: false,
		},
		{
			name:        "large numbers with remainder",
			a:           uint256.MustFromDecimal("1000000000000000001"),
			b:           uint256.MustFromDecimal("1000000000000000001"),
			denominator: uint256.MustFromDecimal("1000000000000000000"),
			expected:    uint256.MustFromDecimal("1000000000000000003"), // rounds up
			expectError: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := new(uint256.Int)
			resultPtr, err := MulDivRoundingUp(result, tc.a, tc.b, tc.denominator)

			if tc.expectError {
				require.Error(t, err)
				require.Equal(t, ErrorDivideByZero, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, resultPtr)
				require.Equal(t, result, resultPtr) // Check that the returned pointer is the same as the input
				require.Equal(t, tc.expected.String(), result.String(), "Expected %s but got %s", tc.expected.String(), result.String())
			}
		})
	}
}
