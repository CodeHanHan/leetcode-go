package goden0409

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_BSTSequences(t *testing.T) {
	tree := BuildTree([]int{5, 3, 2, 1, 4, 6}, []int{1, 2, 3, 4, 5, 6})
	require.Equal(t, [][]int{
		{5, 3, 6, 2, 4, 1},
		{5, 3, 6, 2, 1, 4},
		{5, 3, 6, 4, 2, 1},
		{5, 3, 2, 4, 6, 1},
		{5, 3, 2, 4, 1, 6},
		{5, 3, 2, 6, 1, 4},
		{5, 3, 2, 6, 4, 1},
		{5, 3, 2, 1, 4, 6},
		{5, 3, 2, 1, 6, 4},
		{5, 3, 4, 6, 2, 1},
		{5, 3, 4, 2, 6, 1},
		{5, 3, 4, 2, 1, 6},
		{5, 6, 3, 2, 4, 1},
		{5, 6, 3, 2, 1, 4},
		{5, 6, 3, 4, 2, 1}}, BSTSequences(tree))
}
