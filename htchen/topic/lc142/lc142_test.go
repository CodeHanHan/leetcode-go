package lc142

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/LinkList"
	"github.com/stretchr/testify/require"
)

func Test_detectCycle(t *testing.T) {
	l1 := BuildListWithNoHead([]int{3, 2, 0, -4})
	n := l1.Index(2)
	l1tail := l1.Tail()
	l1tail.Next = n
	require.Equal(t, n, detectCycle(l1))
}
