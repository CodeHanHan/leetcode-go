package lc229

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_majorityElement1(t *testing.T) {
	require.ElementsMatch(t, []int{3}, majorityElement1([]int{3, 2, 3}))
	require.ElementsMatch(t, []int{1}, majorityElement1([]int{1}))
	require.ElementsMatch(t, []int{1, 2}, majorityElement1([]int{1, 1, 1, 3, 3, 2, 2, 2}))

}

func Test_majorityElement2(t *testing.T) {
	require.ElementsMatch(t, []int{3}, majorityElement2([]int{3, 2, 3}))
	require.ElementsMatch(t, []int{1}, majorityElement2([]int{1}))
	require.ElementsMatch(t, []int{1, 2}, majorityElement2([]int{1, 1, 1, 3, 3, 2, 2, 2}))
}
