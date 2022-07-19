package topic704

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_search(t *testing.T) {
	require.Equal(t, 4, search([]int{-1, 0, 3, 5, 9, 12}, 9))

	require.Equal(t, -1, search([]int{-1, 0, 3, 5, 9, 12}, 2))

	require.Equal(t, 0, search([]int{5}, 5))

	require.Equal(t, -1, search([]int{}, 5))
}
