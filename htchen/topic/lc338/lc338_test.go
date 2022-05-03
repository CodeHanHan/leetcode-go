package lc338

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countBits(t *testing.T) {
	require.Equal(t, []int{0, 1, 1}, countBits(2))
	
    require.Equal(t, []int{0, 1, 1, 2, 1, 2}, countBits(5))
}
