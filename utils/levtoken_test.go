package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLevTokenUtils(t *testing.T) {
	require.Equal(t, ToDecimal(MaxSupply, 18).String(), "10000000")
	require.Equal(t, ToDecimal(TeamTokens, 18).String(), "7500")
	require.Equal(t, ToDecimal(TeamTokensVested, 18).String(), "5625")
	require.Equal(t, VestingPeriodMonths, 12)
	require.Equal(t, VestingCycleMonths, 1)
}
