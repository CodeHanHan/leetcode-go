package offer_10_II_s

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numWays(t *testing.T) {
	require.Equal(t, 2, numWays(2))

	require.Equal(t, 21, numWays(7))

	require.Equal(t, 1, numWays(0))
}
