package goden1710

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_majorityElement(t *testing.T) {
	require.Equal(t, 2, majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	require.Equal(t, -1, majorityElement([]int{2, 1}))
	require.Equal(t, 3, majorityElement([]int{3, 2, 3}))
	require.Equal(t, 5, majorityElement([]int{1, 2, 5, 9, 5, 9, 5, 5, 5}))
	require.Equal(t, 5, majorityElement([]int{6, 5, 5}))
	require.Equal(t, -1, majorityElement([]int{1, 2, 3}))
}
