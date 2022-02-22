package topic62

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_uniquePaths(t *testing.T) {
	require.Equal(t, 1, uniquePaths(1, 1))

	require.Equal(t, 1, uniquePaths(1, 1000))

	require.Equal(t, 1, uniquePaths(1000, 1))

	require.Equal(t, 3, uniquePaths(3, 2))

	require.Equal(t, 28, uniquePaths(3, 7))

	require.Equal(t, 28, uniquePaths(7, 3))

	require.Equal(t, 6, uniquePaths(3, 3))
}

func Test_uniquePaths_1(t *testing.T) {
	require.Equal(t, 1, uniquePaths_1(1, 1))

	require.Equal(t, 1, uniquePaths_1(1, 1000))

	require.Equal(t, 1, uniquePaths_1(1000, 1))

	require.Equal(t, 3, uniquePaths_1(3, 2))

	require.Equal(t, 28, uniquePaths_1(3, 7))

	require.Equal(t, 28, uniquePaths_1(7, 3))

	require.Equal(t, 6, uniquePaths_1(3, 3))
}
