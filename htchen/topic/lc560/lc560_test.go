package lc560

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_subarraySum1(t *testing.T) {
	require.Equal(t, 2, subarraySum1([]int{1, 1, 1}, 2))

	require.Equal(t, 2, subarraySum1([]int{1, 2, 3}, 3))
}

func Test_subarraySum2(t *testing.T) {
	require.Equal(t, 2, subarraySum2([]int{1, 1, 1}, 2))

	require.Equal(t, 2, subarraySum2([]int{1, 2, 3}, 3))
}
