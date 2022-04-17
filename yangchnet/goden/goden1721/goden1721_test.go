package goden1721

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_trap(t *testing.T) {
	require.Equal(t, 6, trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	require.Equal(t, 0, trap([]int{0, 1}))
}
