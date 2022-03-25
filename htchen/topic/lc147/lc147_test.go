package lc147

import (
	"testing"

	list "github.com/CodeHanHan/leetcode-go/base/LinkList"
	"github.com/stretchr/testify/require"
)

func Test_insertionSortList(t *testing.T) {
	require.Equal(t, list.BuildListWithNoHead([]int{1, 2, 3, 4}),
		insertionSortList(list.BuildListWithNoHead([]int{4, 2, 1, 3})))

	require.Equal(t, list.BuildListWithNoHead([]int{-1, 0, 3, 4, 5}),
		insertionSortList(list.BuildListWithNoHead([]int{-1, 5, 3, 4, 0})))
}
