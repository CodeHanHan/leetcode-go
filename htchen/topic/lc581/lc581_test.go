package lc581

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findUnsortedSubarray(t *testing.T) {
	require.Equal(t, 5, findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))

	require.Equal(t, 0, findUnsortedSubarray([]int{1, 2, 3, 4}))
}

func Test_findUnsortedSubarray1(t *testing.T) {
	require.Equal(t, 5, findUnsortedSubarray1([]int{2, 6, 4, 8, 10, 9, 15}))
    
	require.Equal(t, 0, findUnsortedSubarray1([]int{1, 2, 3, 4}))
}
