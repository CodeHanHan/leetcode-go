package topic70

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_climbStairs(t *testing.T) {
	require.Equal(t, 1, climbStairs(1))

	require.Equal(t, 2, climbStairs(2))

	require.Equal(t, 3, climbStairs(3))
}

func Test_climbStairs_1(t *testing.T) {
	require.Equal(t, 1, climbStairs_1(1))

	require.Equal(t, 2, climbStairs_1(2))

	require.Equal(t, 3, climbStairs_1(3))
}
