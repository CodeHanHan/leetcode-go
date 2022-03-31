package lc414

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_thirdMax1(t *testing.T) {
	require.Equal(t, 1, thirdMax1([]int{3, 2, 1}))
	
    require.Equal(t, 2, thirdMax1([]int{1, 2}))
	
    require.Equal(t, 1, thirdMax1([]int{2, 2, 3, 1}))
}

func Test_thirdMax2(t *testing.T) {
	require.Equal(t, 1, thirdMax1([]int{3, 2, 1}))
	
    require.Equal(t, 2, thirdMax1([]int{1, 2}))
	
    require.Equal(t, 1, thirdMax1([]int{2, 2, 3, 1}))
}
