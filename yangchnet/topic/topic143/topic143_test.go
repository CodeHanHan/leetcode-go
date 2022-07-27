package topic143

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/LinkList"
	"github.com/stretchr/testify/require"
)

func Test_reorderList(t *testing.T) {
	var list *ListNode
	list = BuildListWithNoHead([]int{1, 2, 3, 4})

	reorderList(list)
	require.Equal(t, "1 → 4 → 2 → 3 → nil", list.String())

	list = BuildListWithNoHead([]int{1, 2, 3, 4, 5})
	reorderList(list)
	require.Equal(t, "1 → 5 → 2 → 4 → 3 → nil", list.String())
}
