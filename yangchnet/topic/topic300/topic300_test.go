package topic300

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_lengthOfLIS(t *testing.T) {
	require.Equal(t, 4, lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	require.Equal(t, 5, lengthOfLIS([]int{1, 2, 3, 4, 5}))
	require.Equal(t, 1, lengthOfLIS([]int{5, 4, 3, 2, 1}))
	require.Equal(t, 1, lengthOfLIS([]int{1}))
	require.Equal(t, 0, lengthOfLIS([]int{}))
}

func Test_lengthOfLIS2(t *testing.T) {
	require.Equal(t, 4, lengthOfLIS2([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	require.Equal(t, 5, lengthOfLIS2([]int{1, 2, 3, 4, 5}))
	require.Equal(t, 1, lengthOfLIS2([]int{5, 4, 3, 2, 1}))
	require.Equal(t, 1, lengthOfLIS2([]int{1}))
	require.Equal(t, 0, lengthOfLIS2([]int{}))
}
