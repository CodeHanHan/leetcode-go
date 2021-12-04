package topic039

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_combinationSum(t *testing.T) {
	require.Equal(t,
		[][]int{{2, 2, 3}, {7}},
		combinationSum([]int{2, 3, 6, 7}, 7))

	require.Equal(t,
		[][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		combinationSum([]int{2, 3, 5}, 8))

	require.Equal(t,
		[][]int{},
		combinationSum([]int{2}, 1))

	require.Equal(t,
		[][]int{{1}},
		combinationSum([]int{1}, 1))

	require.Equal(t,
		[][]int{{1, 1, 1, 1, 1, 1}, {1, 1, 1, 3}, {3, 3}},
		combinationSum([]int{1, 3}, 6))

	require.Equal(t,
		[][]int{{1, 1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 3}, {1, 1, 1, 3, 3}, {3, 3, 3}},
		combinationSum([]int{1, 3}, 9))

	require.Equal(t,
		[][]int{{1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 2},
			{1, 1, 1, 1, 1, 1, 3},
			{1, 1, 1, 1, 1, 2, 2},
			{1, 1, 1, 1, 2, 3},
			{1, 1, 1, 1, 5},
			{1, 1, 1, 2, 2, 2},
			{1, 1, 1, 3, 3},
			{1, 1, 1, 6},
			{1, 1, 2, 2, 3},
			{1, 1, 2, 5},
			{1, 1, 7},
			{1, 2, 2, 2, 2},
			{1, 2, 3, 3},
			{1, 2, 6},
			{1, 3, 5},
			{2, 2, 2, 3},
			{2, 2, 5},
			{2, 7},
			{3, 3, 3},
			{3, 6}},
		combinationSum([]int{2, 7, 6, 3, 5, 1}, 9))
}
