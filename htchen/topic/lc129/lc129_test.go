package lc129

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/Tree"
	"github.com/stretchr/testify/require"
)

func Test_sumNumbers(t *testing.T) {
	require.Equal(t, 25,
		sumNumbers(
			BuildTree([]int{1, 2, 3}, []int{2, 1, 3})))
	require.Equal(t, 1026,
		sumNumbers(
			BuildTree([]int{4, 9, 5, 1, 0}, []int{5, 9, 1, 4, 0})))
}

func Test_sumNumbers1(t *testing.T) {
	require.Equal(t, 25,
		sumNumbers1(
			BuildTree([]int{1, 2, 3}, []int{2, 1, 3})))
	require.Equal(t, 1026,
		sumNumbers1(
			BuildTree([]int{4, 9, 5, 1, 0}, []int{5, 9, 1, 4, 0})))
}
