package lc347

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_topKFrequent(t *testing.T) {
	require.Equal(t, []int{1, 2}, topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))

	require.Equal(t, []int{1}, topKFrequent([]int{1}, 1))
}
