package topic001

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_twoSum(t *testing.T) {
	require.Equal(t, []int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9))

	require.Equal(t, []int{1, 2}, twoSum([]int{3, 2, 4}, 6))

	require.Equal(t, []int{0, 1}, twoSum([]int{3, 3}, 6))

	require.Equal(t, []int{0, 6}, twoSum_1([]int{3, 5, 0, 1, 2, 4, -3}, 0))
}

func Test_twoSum_1(t *testing.T) {
	require.Equal(t, []int{0, 1}, twoSum_1([]int{2, 7, 11, 15}, 9))

	require.Equal(t, []int{1, 2}, twoSum_1([]int{3, 2, 4}, 6))

	require.Equal(t, []int{0, 1}, twoSum_1([]int{3, 3}, 6))

	require.Equal(t, []int{0, 6}, twoSum_1([]int{3, 5, 0, 1, 2, 4, -3}, 0))
}
