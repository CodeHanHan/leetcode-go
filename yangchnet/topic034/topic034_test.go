package topic034

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_searchRange(t *testing.T) {
	require.Equal(t, []int{3, 4}, searchRange([]int{5, 7, 7, 8, 8, 10}, 8))

	require.Equal(t, []int{3, 4}, searchRange([]int{5, 7, 7, 8, 8}, 8))

	require.Equal(t, []int{0, 0}, searchRange([]int{5, 7, 7, 8, 8}, 5))

	require.Equal(t, []int{0, 2}, searchRange([]int{5, 5, 5}, 5))
}

func Test_searchRange_1(t *testing.T) {
	require.Equal(t, []int{3, 4}, searchRange_1([]int{5, 7, 7, 8, 8, 10}, 8))

	require.Equal(t, []int{3, 4}, searchRange_1([]int{5, 7, 7, 8, 8}, 8))

	require.Equal(t, []int{0, 0}, searchRange_1([]int{5, 7, 7, 8, 8}, 5))

	require.Equal(t, []int{0, 2}, searchRange_1([]int{5, 5, 5}, 5))

	require.Equal(t, []int{-1, -1}, searchRange_1([]int{5, 7, 7, 8, 8, 10}, 6))
}

func Test_searchRange_2(t *testing.T) {
	require.Equal(t, []int{3, 4}, searchRange_2([]int{5, 7, 7, 8, 8, 10}, 8))

	require.Equal(t, []int{3, 4}, searchRange_2([]int{5, 7, 7, 8, 8}, 8))

	require.Equal(t, []int{0, 0}, searchRange_2([]int{5, 7, 7, 8, 8}, 5))

	require.Equal(t, []int{0, 2}, searchRange_2([]int{5, 5, 5}, 5))

	require.Equal(t, []int{-1, -1}, searchRange_2([]int{5, 7, 7, 8, 8, 10}, 6))
}
