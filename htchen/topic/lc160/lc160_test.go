package lc160

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/LinkList"
	"github.com/stretchr/testify/require"
)

func Test_getIntersectionNode(t *testing.T) {
	l1 := BuildListWithNoHead([]int{1, 2, 3, 4, 5})
	n := l1.Index(3)
	l2 := BuildListWithNoHead([]int{6, 7})
	l2tail := l2.Tail()
	l2tail.Next = n
	require.Equal(t, n,
		getIntersectionNode(l1, l2))
}
