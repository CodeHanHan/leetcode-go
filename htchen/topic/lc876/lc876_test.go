package lc876

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/LinkList"
	"github.com/stretchr/testify/require"
)

func Test_middleNode(t *testing.T) {
	require.Equal(t,
		Index(BuildListWithNoHead([]int{1, 2, 3, 4, 5}), 3),
		middleNode(BuildListWithNoHead([]int{1, 2, 3, 4, 5})))
	require.Equal(t,
		Index(BuildListWithNoHead([]int{1, 2, 3, 4, 5, 6}), 4),
		middleNode(BuildListWithNoHead([]int{1, 2, 3, 4, 5, 6})))
}

func Index(l *ListNode, target int) *ListNode {
	if l == nil {
		return nil
	}

	p := l
	for p != nil {
		if p.Val == target {
			return p
		}
		p = p.Next
	}

	return nil
}
