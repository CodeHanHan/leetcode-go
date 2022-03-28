package offer24

import (
	"testing"

	"github.com/stretchr/testify/require"

	. "github.com/CodeHanHan/leetcode-go/base/LinkList"
)

func Test_reverseList(t *testing.T) {
	var linklist *ListNode = BuildListWithNoHead([]int{5, 4, 3, 2, 1})
	require.Equal(t, "1 → 2 → 3 → 4 → 5 → nil", reverseList(linklist).String())
}

func Test_reverseList2(t *testing.T) {
	var linklist *ListNode = BuildListWithNoHead([]int{5, 4, 3, 2, 1})
	require.Equal(t, "1 → 2 → 3 → 4 → 5 → nil", reverseList2(linklist).String())
}
