package offer03

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findRepeatNumber(t *testing.T) {

	require.Equal(t, 2, findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3}))

	require.Equal(t, 3, findRepeatNumber([]int{3, 1, 0, 2, 5, 3}))
	
	require.Equal(t, -1, findRepeatNumber([]int{3, 1, 0, 2, 5, 4}))

	
}
