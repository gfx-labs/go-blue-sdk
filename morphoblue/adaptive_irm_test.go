package morphoblue

import (
	"fmt"
	"testing"

	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
)

func TestExpLib(t *testing.T) {
	type testcase struct {
		In  *uint256.Int
		Out *uint256.Int
	}
	testCases := []testcase{
		{
			In:  uint256.NewInt(0).Neg(uint256.MustFromDecimal("1000000000000000000000000")),
			Out: uint256.NewInt(0),
		},
		{
			In:  uint256.NewInt(0).Neg(uint256.MustFromDecimal("1907753029319520")),
			Out: uint256.NewInt(998094066731490918),
		},
		// Positive exponents (q > 0): these returned 0 before the fix.
		{
			In:  uint256.MustFromDecimal("693147180559945309"), // ln 2
			Out: uint256.MustFromDecimal("2000000000000000000"),
		},
		{
			In:  uint256.MustFromDecimal("1000000000000000000"), // e^1
			Out: uint256.MustFromDecimal("2707864291678420188"),
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Case%d", idx), func(t *testing.T) {
			out := AdaptiveIRM.ExpLibWExp(tc.In)
			require.Equal(t, tc.Out.String(), out.String())
		})
	}
}

func TestAdaptiveCurve(t *testing.T) {
	type testcase struct {
		Utilization *uint256.Int
		StartRate   *uint256.Int
		Duration    *uint256.Int

		BorrowRate *uint256.Int
		EndRate    *uint256.Int
	}
	testCases := []testcase{
		{
			Utilization: uint256.NewInt(559455285663861780),
			StartRate:   uint256.NewInt(240238572),
			Duration:    uint256.NewInt(4490),
			BorrowRate:  uint256.NewInt(171830421),
			EndRate:     uint256.NewInt(239592324),
		},
		// Over-target market idle for ~7.5 days (Pharos market 0xba0d26..., block
		// 19628368). BorrowRate equals the on-chain AdaptiveCurveIrm
		// borrowRateView for the same state and timestamp. Before the wExp fix
		// this returned 372143107 with EndRate clamped to MIN_RATE_AT_TARGET.
		{
			Utilization: uint256.NewInt(980898589158542438),
			StartRate:   uint256.NewInt(339242037),
			Duration:    uint256.NewInt(650244),
			BorrowRate:  uint256.NewInt(1845453592),
			EndRate:     uint256.NewInt(780802641),
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Case%d", idx), func(t *testing.T) {
			borrowRate, endRate := AdaptiveIRM.GetBorrowRate(
				tc.Utilization,
				tc.StartRate,
				tc.Duration,
			)
			require.Equal(t, tc.BorrowRate.String(), borrowRate.String())
			require.Equal(t, tc.EndRate.String(), endRate.String())
		})
	}
}
