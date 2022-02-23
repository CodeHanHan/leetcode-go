package lc268

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_missingNumber(t *testing.T) {

	require.Equal(t, 2, missingNumber([]int{3, 0, 1}))

	require.Equal(t, 2, missingNumber([]int{0, 1}))

	require.Equal(t, 8, missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))

	require.Equal(t, 1, missingNumber([]int{0}))

	require.Equal(t, 2, missingNumber([]int{3, 0, 1, 4, 5}))
}

func Test_missingNumber_1(t *testing.T) {

	require.Equal(t, 2, missingNumber([]int{3, 0, 1}))

	require.Equal(t, 2, missingNumber([]int{0, 1}))

	require.Equal(t, 8, missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))

	require.Equal(t, 1, missingNumber([]int{0}))

	require.Equal(t, 2, missingNumber([]int{3, 0, 1, 4, 5}))
}

func Test_missingNumber_2(t *testing.T) {

	require.Equal(t, 2, missingNumber([]int{3, 0, 1}))

	require.Equal(t, 2, missingNumber([]int{0, 1}))

	require.Equal(t, 8, missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))

	require.Equal(t, 1, missingNumber([]int{0}))

	require.Equal(t, 2, missingNumber([]int{3, 0, 1, 4, 5}))

	require.Equal(t, 4, missingNumber([]int{3, 2, 0, 1, 5}))
}
