package goden1607

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximum(t *testing.T) {
	require.Equal(t, 2, maximum(1, 2))
	require.Equal(t, 5, maximum(2, 5))
	require.Equal(t, 2147483647, maximum(2147483647, -2147483648))
}
