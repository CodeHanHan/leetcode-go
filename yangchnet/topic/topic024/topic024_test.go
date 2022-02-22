package topic024

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_swapPairs(t *testing.T) {
	require.Equal(t, BuildListWithNoHead([]int{2, 1, 4, 3}), swapPairs(BuildListWithNoHead([]int{1, 2, 3, 4})))

	require.Equal(t, BuildListWithNoHead([]int{2, 1, 4, 3, 5}), swapPairs(BuildListWithNoHead([]int{1, 2, 3, 4, 5})))

	require.Equal(t, BuildListWithNoHead([]int{1}), swapPairs(BuildListWithNoHead([]int{1})))

	require.Equal(t, BuildListWithNoHead([]int{}), swapPairs(BuildListWithNoHead([]int{})))
}
