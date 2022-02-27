package lc094

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_inorderTraversal(t *testing.T) {
	require.Equal(t, []int{1, 3, 2}, inorderTraversal(BuildTree([]int{1, 2, 3}, []int{1, 3, 2})))

	require.Equal(t, []int{1}, inorderTraversal(BuildTree([]int{1}, []int{1})))

	require.Equal(t, []int{2, 1}, inorderTraversal(BuildTree([]int{1, 2}, []int{2, 1})))
}
