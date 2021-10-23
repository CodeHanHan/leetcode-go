package LC001

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_twoSum(t *testing.T) {
	require.Equal(t, []int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9))

	require.Equal(t, []int{1, 2}, twoSum([]int{3, 2, 4}, 6))

	require.Equal(t, []int{0, 1}, twoSum([]int{3, 3}, 6))

}
