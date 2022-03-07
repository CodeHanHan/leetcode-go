package goden0404

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isBalanced(t *testing.T) {
	tree := BuildTree([]int{1}, []int{1})
	require.Equal(t, true, isBalanced(tree))
}
