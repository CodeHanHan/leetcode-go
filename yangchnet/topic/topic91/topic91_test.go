package topic91

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numDecodings(t *testing.T) {
	require.Equal(t, 2, numDecodings("12"))

	require.Equal(t, 3, numDecodings("226"))

	require.Equal(t, 0, numDecodings("0"))

	require.Equal(t, 0, numDecodings("2015016"))

	require.Equal(t, 1, numDecodings("10"))
}
