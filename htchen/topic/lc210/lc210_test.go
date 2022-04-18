package lc210

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findOrder(t *testing.T) {
	require.ElementsMatch(t, []int{0, 1},
		findOrder(2, [][]int{{1, 0}}))

	require.ElementsMatch(t, []int{0, 2, 1, 3},
		findOrder(4, [][]int{{1, 0}, []int{2, 0}, []int{3, 1}, []int{3, 2}}))
}
