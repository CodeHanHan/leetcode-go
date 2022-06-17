package topic47

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_permuteUnique(t *testing.T) {

	// t.Log(permuteUnique([]int{1, 2, 2}))

	require.ElementsMatch(t, [][]int{{1, 2, 2}, {2, 2, 1}, {2, 1, 2}}, permuteUnique([]int{1, 2, 2}))

	require.ElementsMatch(t, [][]int{{1, 2}, {2, 1}}, permuteUnique([]int{1, 2}))

	require.ElementsMatch(t, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}, permuteUnique([]int{1, 2, 3}))

	require.ElementsMatch(t, [][]int{{1}}, permuteUnique([]int{1}))
}

func Test_permuteUnique1(t *testing.T) {
	require.ElementsMatch(t, [][]int{{1, 2, 2}, {2, 2, 1}, {2, 1, 2}}, permuteUnique1([]int{1, 2, 2}))

	require.ElementsMatch(t, [][]int{{1, 2}, {2, 1}}, permuteUnique1([]int{1, 2}))

	require.ElementsMatch(t, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}, permuteUnique1([]int{1, 2, 3}))

	require.ElementsMatch(t, [][]int{{1}}, permuteUnique1([]int{1}))
}
