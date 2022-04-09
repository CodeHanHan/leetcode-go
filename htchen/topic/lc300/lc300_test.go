package lc300

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_lengthOfLIS(t *testing.T) {
	require.Equal(t, 4, lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	
    require.Equal(t, 4, lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	
    require.Equal(t, 1, lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
}
