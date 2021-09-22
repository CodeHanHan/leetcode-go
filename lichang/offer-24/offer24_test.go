package offer24

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_traverseList(t *testing.T) {
	require.Equal(t, []int{1, 2, 3, 4, 5}, traverseList(linklist_5))
}

func Test_reverseList(t *testing.T) {
	require.Equal(t, []int{5, 4, 3, 2, 1}, traverseList(reverseList(linklist_5)))
}

var linklist_5 = buildList([]int{1, 2, 3, 4, 5})

func buildList(nums []int) *ListNode {
	if len(nums) < 1 {
		return nil
	}

	head := &ListNode{
		Val:  nums[0],
		Next: nil,
	}

	for i := 1; i < len(nums); i++ {
		insertToTail(head, nums[i])
	}
	return head
}
