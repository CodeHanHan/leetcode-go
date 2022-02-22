package topic142

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_detectCycle(t *testing.T) {
	require.Equal(t, 2, detectCycle(
		BuildCircleListWithNoHead([]int{3, 2, 0, -4}, 1),
	).Val)

	require.Equal(t, 1, detectCycle(
		BuildCircleListWithNoHead([]int{1, 2}, 0),
	).Val)

	require.Equal(t, 1, detectCycle(
		BuildCircleListWithNoHead([]int{1}, 0),
	).Val)

	require.Equal(t, (*ListNode)(nil), detectCycle(
		BuildListWithNoHead([]int{1}),
	))
}
