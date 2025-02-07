package morphoblue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxValues(t *testing.T) {
	require.Equal(t, MaxUint256.String(), "115792089237316195423570985008687907853269984665640564039457584007913129639935")
	require.Equal(t, MaxUint160.String(), "1461501637330902918203684832716283019655932542975")
	require.Equal(t, MaxUint128.String(), "340282366920938463463374607431768211455")
}
