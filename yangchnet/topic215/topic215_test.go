package topic215

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findKthLargest(t *testing.T) {
	require.Equal(t, 5, findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))

	require.Equal(t, 4, findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}

func Test_findKthLargest_2(t *testing.T) {
	require.Equal(t, 5, findKthLargest_2([]int{3, 2, 1, 5, 6, 4}, 2))

	require.Equal(t, 4, findKthLargest_2([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}

func Test_findKthLargest_3(t *testing.T) {
	require.Equal(t, 5, findKthLargest_3([]int{3, 2, 1, 5, 6, 4}, 2))

	require.Equal(t, 4, findKthLargest_3([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
