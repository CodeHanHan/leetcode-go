package lc024

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_swapPairs(t *testing.T) {
	require.Equal(t,
		BuildListWithNoHead([]int{1, 2, 3, 4}),
		swapPairs(BuildListWithNoHead([]int{2, 1, 4, 3})))

	require.Equal(t,
		BuildListWithNoHead([]int{}),
		swapPairs(BuildListWithNoHead([]int{})))

	require.Equal(t,
		BuildListWithNoHead([]int{1}),
		swapPairs(BuildListWithNoHead([]int{1})))
}
