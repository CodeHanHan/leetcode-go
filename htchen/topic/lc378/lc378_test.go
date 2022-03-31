package lc378

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_kthSmallest1(t *testing.T) {
	require.Equal(t, 13, kthSmallest1([][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8))

	require.Equal(t, -5, kthSmallest1([][]int{{-5}}, 1))
}

func Test_kthSmallest2(t *testing.T) {
	require.Equal(t, 13, kthSmallest2([][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8))

	require.Equal(t, -5, kthSmallest2([][]int{{-5}}, 1))
}
