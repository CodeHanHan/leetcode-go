package lc062

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_uniquePaths(t *testing.T) {
	require.Equal(t, 28, uniquePaths(3, 7))
    
	require.Equal(t, 28, uniquePaths(7, 3))

	require.Equal(t, 3, uniquePaths(3, 2))
}
