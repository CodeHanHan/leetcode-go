package topic41

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_firstMissingPositive(t *testing.T) {
	require.Equal(t, 3, firstMissingPositive([]int{1, 2, 0}))

	require.Equal(t, 2, firstMissingPositive([]int{3, 4, -1, 1}))

	require.Equal(t, 1, firstMissingPositive([]int{7, 8, 9, 11, 12}))
}

func Test_firstMissingPositive1(t *testing.T) {
	// require.Equal(t, 3, firstMissingPositive1([]int{1, 2, 0}))

	require.Equal(t, 2, firstMissingPositive1([]int{3, 4, -1, 1}))

	// require.Equal(t, 1, firstMissingPositive1([]int{7, 8, 9, 11, 12}))
}
