package lc066

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_plusOne(t *testing.T) {

	require.Equal(t, []int{1, 2, 4}, plusOne([]int{1, 2, 3}))

	require.Equal(t, []int{4, 3, 2, 2}, plusOne([]int{4, 3, 2, 1}))

	require.Equal(t, []int{1}, plusOne([]int{0}))
}
