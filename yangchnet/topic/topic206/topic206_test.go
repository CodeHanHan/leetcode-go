package topic206

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/LinkList"
	"github.com/stretchr/testify/require"
)

func Test_reverseList(t *testing.T) {
	require.Equal(t, "5 → 4 → 3 → 2 → 1 → nil", reverseList(BuildListWithNoHead([]int{1, 2, 3, 4, 5})).String())
}

func Test_reverseList1(t *testing.T) {
	require.Equal(t, "5 → 4 → 3 → 2 → 1 → nil", reverseList1(BuildListWithNoHead([]int{1, 2, 3, 4, 5})).String())
}
