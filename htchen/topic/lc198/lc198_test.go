package lc198

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_rob(t *testing.T) {
	require.Equal(t, 4, rob([]int{1, 2, 3, 1}))
	
    require.Equal(t, 12, rob([]int{2, 7, 9, 3, 1}))
}
