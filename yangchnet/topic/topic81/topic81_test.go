package topic81

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_search(t *testing.T) {
	require.Equal(t, true, search([]int{4, 5, 6, 7, 0, 1, 2}, 0))

	require.Equal(t, false, search([]int{4, 5, 6, 7, 0, 1, 2}, 3))

	require.Equal(t, false, search([]int{4, 4, 5, 5, 6, 7, 0, 0, 1, 2}, 3))

	require.Equal(t, false, search([]int{1}, 0))

	require.Equal(t, false, search([]int{}, 0))

	require.Equal(t, true, search([]int{1, 0, 1, 1, 1}, 0))
}
