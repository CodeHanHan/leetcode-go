package topic66

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_plusOne(t *testing.T) {
	require.Equal(t, []int{4, 3, 2, 2}, plusOne([]int{4, 3, 2, 1}))

	require.Equal(t, []int{4, 4, 0}, plusOne([]int{4, 3, 9}))

	require.Equal(t, []int{1, 0}, plusOne([]int{9}))

	require.Equal(t, []int{1}, plusOne([]int{0}))
}
