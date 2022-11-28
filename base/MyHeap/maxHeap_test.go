package heap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_buildHeap(t *testing.T) {
	nums := []int{0, 1, 7, 3, 5, 4, 2, 8}

	buildMaxHeap(nums, 7)

	require.Equal(t, nums, []int{0, 8, 7, 3, 5, 4, 2, 1})
}

func TestMaxSot(t *testing.T) {
	nums := []int{0, 1, 7, 3, 5, 4, 2, 8}
	sortMax(nums, 7)

	require.Equal(t, []int{0, 1, 2, 3, 4, 5, 7, 8}, nums)
}
