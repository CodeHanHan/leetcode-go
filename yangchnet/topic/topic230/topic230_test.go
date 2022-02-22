package topic230

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_kthSmallest(t *testing.T) {
	require.Equal(t, 1,
		kthSmallest(BuildTree([]int{3, 1, 2, 4}, []int{1, 2, 3, 4}),
			1,
		))

	require.Equal(t, 3,
		kthSmallest(BuildTree([]int{5, 3, 2, 1, 4, 6}, []int{1, 2, 3, 4, 5, 6}),
			3,
		))

	require.Equal(t, 1,
		kthSmallest(BuildTree([]int{1}, []int{1}),
			1,
		))
}
