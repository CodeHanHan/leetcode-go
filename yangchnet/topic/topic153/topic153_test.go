package topic153

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMin(t *testing.T) {
	require.Equal(t, 1, findMin([]int{4, 5, 6, 1, 2, 3}))
	require.Equal(t, 1, findMin([]int{1, 2}))
	require.Equal(t, 2, findMin([]int{3, 2}))
	require.Equal(t, 1, findMin([]int{1}))
}
