package lc349

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_intersection(t *testing.T) {
	require.ElementsMatch(t, []int{2}, intersection([]int{1, 2, 2, 1}, []int{2, 2}))

	require.ElementsMatch(t, []int{9, 4}, intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
