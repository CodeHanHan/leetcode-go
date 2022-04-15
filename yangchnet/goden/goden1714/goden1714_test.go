package goden1714

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_smallestK(t *testing.T) {
	require.ElementsMatch(t, []int{1, 2, 3, 4}, smallestK([]int{1, 2, 3, 4, 5, 6, 7}, 4))
	require.ElementsMatch(t, []int{0, 1, 1, 2}, smallestK([]int{1, 2, 3, 4, 5, 6, 7, 1, 2, 0}, 4))
	require.ElementsMatch(t, []int{-11, -2, 0, 1}, smallestK([]int{1, 2, 3, 4, 5, 6, 7, -11, -2, 0}, 4))
}

func Test_smallestK2(t *testing.T) {
	require.ElementsMatch(t, []int{1, 2, 3, 4}, smallestK2([]int{1, 2, 3, 4, 5, 6, 7}, 4))
	require.ElementsMatch(t, []int{0, 1, 1, 2}, smallestK2([]int{1, 2, 3, 4, 5, 6, 7, 1, 2, 0}, 4))
	require.ElementsMatch(t, []int{-11, -2, 0, 1}, smallestK2([]int{1, 2, 3, 4, 5, 6, 7, -11, -2, 0}, 4))
}
