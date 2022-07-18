package lc143

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/LinkList"
	"github.com/stretchr/testify/require"
)

func Test_reorderList(t *testing.T) {
	list := BuildListWithNoHead([]int{1, 2, 3, 4})
	reorderList(list)
	require.Equal(t,
		BuildListWithNoHead([]int{1, 4, 2, 3}), list)
}
func Test_reorderList1(t *testing.T) {
	list := BuildListWithNoHead([]int{1, 2, 3, 4})
	reorderList1(list)
	require.Equal(t,
		BuildListWithNoHead([]int{1, 4, 2, 3}), list)
}
