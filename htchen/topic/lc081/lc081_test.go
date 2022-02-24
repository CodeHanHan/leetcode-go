package lc081

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_search(t *testing.T) {

	require.Equal(t, true, search([]int{2, 5, 6, 0, 0, 1, 2}, 0))

	require.Equal(t, false, search([]int{2, 5, 6, 0, 0, 1, 2}, 3))

	require.Equal(t, true, search([]int{3, 5, 6, 3, 3, 3, 3, 3}, 6))
}
