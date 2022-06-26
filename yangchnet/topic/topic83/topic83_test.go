package topic83

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/LinkList"
	"github.com/stretchr/testify/require"
)

func Test_deleteDuplicates(t *testing.T) {
	head := BuildListWithNoHead([]int{1, 2, 3, 3, 4, 4, 5})
	head = deleteDuplicates(head)
	require.Equal(t, "1 → 2 → 3 → 4 → 5 → nil", head.String())

	head1 := BuildListWithNoHead([]int{1, 1, 1, 2, 3})
	head1 = deleteDuplicates(head1)
	require.Equal(t, "1 → 2 → 3 → nil", head1.String())

	head2 := BuildListWithNoHead([]int{1, 1, 1})
	head2 = deleteDuplicates(head2)
	require.Equal(t, "1 → nil", head2.String())

	head3 := BuildListWithNoHead([]int{})
	head3 = deleteDuplicates(head3)
	require.Equal(t, "nil", head3.String())
}
