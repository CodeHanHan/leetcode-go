package offer25

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_mergeTwoLists(t *testing.T) {
	require.Equal(t,
		"1 → 2 → 3 → 4 → 5 → 6 → 7 → 8 → 9 → 10 → nil",
		fmt.Sprintf("%s", mergeTwoLists(linklist1, linklist2)))

	require.Equal(t,
		"1 → 2 → 3 → nil",
		fmt.Sprintf("%s", mergeTwoLists(buildList([]int{1, 2, 3}), buildList([]int{}))))
}

var linklist1 = buildList([]int{1, 3, 5, 7, 9})

var linklist2 = buildList([]int{2, 4, 6, 8, 10})

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
