package lc169

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_majorityElement(t *testing.T) {
	require.Equal(t, 3, majorityElement([]int{3, 2, 3}))
	require.Equal(t, -1, majorityElement([]int{-1, 2, 3, -1, -1}))
	require.Equal(t, 0, majorityElement([]int{0, 0, 3, 2, 3, 0, 0}))
}

func Test_majorityElement_1(t *testing.T) {
	require.Equal(t, 3, majorityElement([]int{3, 2, 3}))
	require.Equal(t, -1, majorityElement([]int{-1, 2, 3, -1, -1}))
	require.Equal(t, 0, majorityElement([]int{0, 0, 3, 2, 3, 0, 0}))
}
