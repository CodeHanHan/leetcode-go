package topic63

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_uniquePathsWithObstacles(t *testing.T) {
	require.Equal(t, 2, uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))

	require.Equal(t, 1, uniquePathsWithObstacles([][]int{{0, 1}, {0, 0}}))
}
