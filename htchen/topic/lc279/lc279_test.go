package lc279

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numSquares1(t *testing.T) {
	require.Equal(t, 3, numSquares1(12))

	require.Equal(t, 2, numSquares1(13))
}

func Test_numSquares2(t *testing.T) {
	require.Equal(t, 3, numSquares2(12))

	require.Equal(t, 2, numSquares2(13))
}
