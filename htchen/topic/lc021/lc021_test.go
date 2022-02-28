package lc021

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_mergeTwoLists(t *testing.T) {

	require.Equal(t,
		BuildListWithNoHead([]int{1, 1, 2, 3, 4, 4}),
		mergeTwoLists(
			BuildListWithNoHead([]int{1, 2, 4}),
			BuildListWithNoHead([]int{1, 3, 4}),
		),
	)

	require.Equal(t,
		BuildListWithNoHead([]int{}),
		mergeTwoLists(
			BuildListWithNoHead([]int{}),
			BuildListWithNoHead([]int{}),
		),
	)

	require.Equal(t,
		BuildListWithNoHead([]int{0}),
		mergeTwoLists(
			BuildListWithNoHead([]int{}),
			BuildListWithNoHead([]int{0}),
		),
	)
}
