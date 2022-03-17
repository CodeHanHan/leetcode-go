package lc045

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_jump(t *testing.T) {
	require.Equal(t, 2, jump([]int{2, 3, 1, 1, 4}))
    
	require.Equal(t, 2, jump([]int{2, 3, 0, 1, 4}))
}
