package dynamic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Coin(t *testing.T) {
	require.Equal(t, 3, CoinChange(9, []int{1, 3, 5}))
	require.Equal(t, 1, CoinChange(1, []int{1, 3, 5}))
	require.Equal(t, 2, CoinChange(10, []int{1, 3, 5}))
}
