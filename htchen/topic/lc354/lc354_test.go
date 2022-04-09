package lc354

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxEnvelopes(t *testing.T) {
	require.Equal(t, 3,
		maxEnvelopes([][]int{[]int{5, 4}, []int{6, 4}, []int{6, 7}, []int{2, 3}}))

	require.Equal(t, 1,
		maxEnvelopes([][]int{[]int{1, 1}, []int{1, 1}, []int{1, 1}}))
}
