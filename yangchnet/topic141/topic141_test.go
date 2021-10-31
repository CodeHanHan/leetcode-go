package topic141

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_hasCycle(t *testing.T) {
	require.Equal(t, true, hasCycle(
		BuildCircleListWithNoHead(
			[]int{1, 2, 3, 4, 5}, 2),
	))

	require.Equal(t, false, hasCycle(
		BuildListWithNoHead(
			[]int{1, 2, 3, 4, 5}),
	))

	require.Equal(t, true, hasCycle(
		BuildCircleListWithNoHead(
			[]int{1}, 0),
	))

	require.Equal(t, false, hasCycle(
		BuildListWithNoHead(
			[]int{1}),
	))

	require.Equal(t, true, hasCycle(
		BuildCircleListWithNoHead(
			[]int{1, 2, 3, 4, 5, 6}, 0),
	))

	require.Equal(t, true, hasCycle(
		BuildCircleListWithNoHead(
			[]int{1, 2, 3, 4, 5, 6}, 5),
	))

	require.Equal(t, false, hasCycle(
		BuildListWithNoHead(
			[]int{}),
	))
}
