package goden0403

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_listOfDepth(t *testing.T) {
	tree := BuildTree([]int{1, 2, 4, 8, 5, 3, 7}, []int{8, 4, 2, 5, 1, 3, 7})
	list := listOfDepth(tree)
	require.Equal(t, []*ListNode{
		BuildListWithNoHead([]int{1}),
		BuildListWithNoHead([]int{2, 3}),
		BuildListWithNoHead([]int{4, 5, 7}),
		BuildListWithNoHead([]int{8}),
	}, list)

	tree = BuildTree([]int{1}, []int{1})
	list = listOfDepth(tree)
	require.Equal(t, []*ListNode{
		BuildListWithNoHead([]int{1}),
	}, list)
}
