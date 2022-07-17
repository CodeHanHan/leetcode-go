package lc637

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"
	"github.com/stretchr/testify/require"
)

func Test_averageOfLevels(t *testing.T) {
	require.Equal(t,[]float64{3.00000, 14.50000, 11.00000},
		averageOfLevels(
			BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})))
}
