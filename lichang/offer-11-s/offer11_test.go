package offer_11_s

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinArray(t *testing.T) {
	nums1 := []int{5, 1, 2, 3, 4}
	require.Equal(t, 1, minArray(nums1))

	nums2 := []int{3, 0, 1, 2}
	require.Equal(t, 0, minArray(nums2))

	nums3 := []int{2, 2, 2, 0, 1}
	require.Equal(t, 0, minArray(nums3))

	nums4 := []int{3, 4, 5, 1, 2}
	require.Equal(t, 1, minArray(nums4))

	nums5 := []int{1, 0}
	require.Equal(t, 0, minArray(nums5))
}
