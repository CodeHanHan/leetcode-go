package lc300

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_lengthOfLIS1(t *testing.T) {
	require.Equal(t, 4, lengthOfLIS1([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	
    require.Equal(t, 4, lengthOfLIS1([]int{0, 1, 0, 3, 2, 3}))
	
    require.Equal(t, 1, lengthOfLIS1([]int{7, 7, 7, 7, 7, 7, 7}))
}

func Test_lengthOfLIS2(t *testing.T) {
	require.Equal(t, 4, lengthOfLIS2([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	
    require.Equal(t, 4, lengthOfLIS2([]int{0, 1, 0, 3, 2, 3}))
	
    require.Equal(t, 1, lengthOfLIS2([]int{7, 7, 7, 7, 7, 7, 7}))
}
