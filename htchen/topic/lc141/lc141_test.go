package lc141

import (
	"testing"

	list "github.com/CodeHanHan/leetcode-go/base/LinkList"
	"github.com/stretchr/testify/require"
)

func Test_hasCycle(t *testing.T) {
	require.Equal(t, true, hasCycle(list.BuildCircleListWithNoHead([]int{3, 2, 0, -4}, 1)))

	require.Equal(t, true, hasCycle(list.BuildCircleListWithNoHead([]int{1, 2}, 0)))
}
