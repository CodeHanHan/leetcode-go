package lc215

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findKthLargest1(t *testing.T) {
	require.Equal(t, 5, findKthLargest1([]int{3, 2, 1, 5, 6, 4}, 2))
	
    require.Equal(t, 4, findKthLargest1([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
func Test_findKthLargest2(t *testing.T) {
	require.Equal(t, 5, findKthLargest2([]int{3, 2, 1, 5, 6, 4}, 2))
	
    require.Equal(t, 4, findKthLargest2([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
