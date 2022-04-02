package lc226

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"
	"github.com/stretchr/testify/require"
)

func Test_invertTree(t *testing.T) {
	require.Equal(t, BuildTree([]int{4, 7, 9, 6, 2, 3, 1}, []int{9, 7, 6, 4, 3, 2, 1}),
		invertTree(BuildTree([]int{4, 2, 1, 3, 7, 6, 9}, []int{1, 2, 3, 4, 6, 7, 9})))
}
