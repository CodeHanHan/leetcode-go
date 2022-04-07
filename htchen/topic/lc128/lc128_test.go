package lc128

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestConsecutive(t *testing.T) {
	require.Equal(t, 4, longestConsecutive([]int{100, 4, 200, 1, 3, 2}))

	require.Equal(t, 9, longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}
