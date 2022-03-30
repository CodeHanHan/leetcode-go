package lc274

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_hIndex1(t *testing.T) {
	require.Equal(t, 3, hIndex1([]int{3, 0, 6, 1, 5}))

	require.Equal(t, 1, hIndex1([]int{1, 3, 1}))
}

func Test_hIndex2(t *testing.T) {
	require.Equal(t, 3, hIndex2([]int{3, 0, 6, 1, 5}))

	require.Equal(t, 1, hIndex2([]int{1, 3, 1}))
}
