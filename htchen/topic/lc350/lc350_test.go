package lc350

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_intersect1(t *testing.T) {
	require.ElementsMatch(t, []int{2, 2}, intersect1([]int{1, 2, 2, 1}, []int{2, 2}))

	require.ElementsMatch(t, []int{4, 9}, intersect1([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}

func Test_intersect2(t *testing.T) {
	require.ElementsMatch(t, []int{2, 2}, intersect2([]int{1, 2, 2, 1}, []int{2, 2}))

	require.ElementsMatch(t, []int{4, 9}, intersect2([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
