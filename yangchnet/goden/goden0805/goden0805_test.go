package goden0805

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_multiply(t *testing.T) {
	require.Equal(t, 0, multiply(0, 1))
	require.Equal(t, 30, multiply(5, 6))
}
