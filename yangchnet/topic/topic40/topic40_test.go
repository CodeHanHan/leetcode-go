package topic40

import (
	"testing"
)

func Test_combinationSum2(t *testing.T) {

	t.Log(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 10))
	// require.ElementsMatch(t, [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
	// 	combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
