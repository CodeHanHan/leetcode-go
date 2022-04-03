package goden1605

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_trailingZeroes(t *testing.T) {
	require.Equal(t, 0, trailingZeroes(3))
	require.Equal(t, 1, trailingZeroes(5))
}
