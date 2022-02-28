package lc023

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_mergeKLists1(t *testing.T) {
	require.Equal(t,
        BuildListWithNoHead([]int{1,1,2,3,4,4,5,6}),
            mergeKLists1(
                []*ListNode{
                    BuildListWithNoHead([]int{1,4,5}),
                    BuildListWithNoHead([]int{1,3,4}),
                    BuildListWithNoHead([]int{2,6})},
            ),
    )

    require.Equal(t,
        BuildListWithNoHead([]int{}),
            mergeKLists1([]*ListNode{BuildListWithNoHead([]int{})}))
}

func Test_mergeKLists2(t *testing.T) {
	require.Equal(t,
        BuildListWithNoHead([]int{1,1,2,3,4,4,5,6}),
            mergeKLists1(
                []*ListNode{
                    BuildListWithNoHead([]int{1,4,5}),
                    BuildListWithNoHead([]int{1,3,4}),
                    BuildListWithNoHead([]int{2,6})},
            ),
    )

    require.Equal(t,
        BuildListWithNoHead([]int{}),
            mergeKLists1([]*ListNode{BuildListWithNoHead([]int{})}))
}