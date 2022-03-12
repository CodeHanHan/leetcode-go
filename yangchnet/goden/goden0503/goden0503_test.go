package goden0503

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverseBits(t *testing.T) {
	require.Equal(t, 8, reverseBits(1775))
	require.Equal(t, 4, reverseBits(7))
	require.Equal(t, 32, reverseBits(-1))
	require.Equal(t, 32, reverseBits(2147483647))
}
