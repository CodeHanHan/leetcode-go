package lc016

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_threeSumClosest(t *testing.T) {
	require.Equal(t, 2, threeSumClosest([]int{-1, 2, 1, -4}, 1))

	require.Equal(t, 0, threeSumClosest([]int{0, 0, 0}, 1))
}
