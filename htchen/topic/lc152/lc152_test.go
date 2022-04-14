package lc152

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxProduct(t *testing.T) {
	require.Equal(t, 6, maxProduct([]int{2, 3, -2, 4}))
	
    require.Equal(t, 0, maxProduct([]int{-2, 0, -1}))
}
