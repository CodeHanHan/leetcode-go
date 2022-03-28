package topic121

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxProfit(t *testing.T) {
	require.Equal(t, 0, maxProfit([]int{7}))

	require.Equal(t, 4, maxProfit([]int{7, 1, 5}))

	require.Equal(t, 0, maxProfit([]int{7, 6, 5, 4, 3, 2, 1}))

	require.Equal(t, 5, maxProfit([]int{7, 1, 5, 3, 6, 4}))

	require.Equal(t, 6, maxProfit([]int{1, 7}))

	require.Equal(t, 0, maxProfit([]int{1, 1, 1, 1, 1}))

	require.Equal(t, 1, maxProfit([]int{4, 1, 2}))
}
