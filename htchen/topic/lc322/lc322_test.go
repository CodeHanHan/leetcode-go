package lc322

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_coinChange(t *testing.T) {
	require.Equal(t, 3, coinChange([]int{1, 2, 5}, 11))

	require.Equal(t, -1, coinChange([]int{2}, 3))
}
