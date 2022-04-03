package goden1606

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_smallestDifference(t *testing.T) {
	require.Equal(t, 3, smallestDifference([]int{1, 3, 15, 11, 2}, []int{23, 127, 235, 19, 8}))
	require.Equal(t, 0, smallestDifference([]int{1}, []int{1}))
}
