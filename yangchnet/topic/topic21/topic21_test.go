package topic21

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
		BuildListWithNoHead([]int{1, 3, 4}),
		mergeTwoLists(
			BuildListWithNoHead([]int{}),
			BuildListWithNoHead([]int{1, 3, 4}),
		),
	)

	require.Equal(t,
		BuildListWithNoHead([]int{1, 1, 2, 2}),
		mergeTwoLists(
			BuildListWithNoHead([]int{1, 2}),
			BuildListWithNoHead([]int{1, 2}),
		),
	)
}
