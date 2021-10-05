package offer_10_I_s

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_fib(t *testing.T) {
	require.Equal(t, 1, fib(1))

	require.Equal(t, 5, fib(5))

	require.Equal(t, 94208912, fib(99))
}
