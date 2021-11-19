package lc033

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_search(t *testing.T) {

	require.Equal(t, 4, search([]int{4, 5, 6, 7, 0, 1, 2}, 0))

	require.Equal(t, -1, search([]int{4, 5, 6, 7, 0, 1, 2}, 3))

	require.Equal(t, -1, search([]int{1}, 0))
}
