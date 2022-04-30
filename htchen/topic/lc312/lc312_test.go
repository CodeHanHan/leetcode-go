package lc312

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxCoins(t *testing.T) {
	require.Equal(t, 167, maxCoins([]int{3, 1, 5, 8}))
	
    require.Equal(t, 10, maxCoins([]int{1, 5}))
}
