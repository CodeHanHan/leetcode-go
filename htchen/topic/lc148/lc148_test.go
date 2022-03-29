package lc148

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/LinkList"
	"github.com/stretchr/testify/require"
)

func Test_sortList(t *testing.T) {
	require.Equal(t, BuildListWithNoHead([]int{1, 2, 3, 4}),
		sortList(BuildListWithNoHead([]int{4, 2, 1, 3})))

	require.Equal(t, BuildListWithNoHead([]int{-1, 0, 3, 4, 5}),
		sortList(BuildListWithNoHead([]int{-1, 5, 3, 4, 0})))
}
