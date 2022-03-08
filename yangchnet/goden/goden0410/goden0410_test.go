package goden0410

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkSubTree(t *testing.T) {
	tree := BuildTree([]int{2, 1, 3}, []int{1, 2, 3})
	n := tree.Find(3)
	require.Equal(t, true, checkSubTree(tree, n))
}
