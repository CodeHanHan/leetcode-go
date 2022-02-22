package topic7

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverse(t *testing.T) {
	require.Equal(t, 123, reverse(321))
	require.Equal(t, 0, reverse(0))
	require.Equal(t, -321, reverse(-123))
	require.Equal(t, 0, reverse(-999999999999999999))
	require.Equal(t, 21, reverse(1200000))
}

func Test_reverse_1(t *testing.T) {
	require.Equal(t, 123, reverse_1(321))
	require.Equal(t, 0, reverse_1(0))
	require.Equal(t, -321, reverse_1(-123))
	require.Equal(t, 0, reverse_1(-999999999999999999))
	require.Equal(t, 21, reverse_1(1200000))
}
