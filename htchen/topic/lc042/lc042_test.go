package lc042

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_trap(t *testing.T) {
	require.Equal(t, 6, trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	
    require.Equal(t, 9, trap([]int{4, 2, 0, 3, 2, 5}))
}

func Test_trap1(t *testing.T) {
	require.Equal(t, 6, trap1([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	
    require.Equal(t, 9, trap1([]int{4, 2, 0, 3, 2, 5}))
}