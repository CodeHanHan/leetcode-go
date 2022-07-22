package lc221

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximalSquare(t *testing.T) {
	require.Equal(t, 4,
		maximalSquare([][]byte{
			[]byte{'1', '0', '1', '0', '0'},
			[]byte{'1', '0', '1', '1', '1'},
			[]byte{'1', '1', '1', '1', '1'},
			[]byte{'1', '0', '0', '1', '0'},
		}))

	require.Equal(t, 1,
		maximalSquare([][]byte{
			[]byte{'0', '1'},
			[]byte{'1', '0'},
		}))
}
func Test_maximalSquare1(t *testing.T) {
	require.Equal(t, 4,
		maximalSquare1([][]byte{
			[]byte{'1', '0', '1', '0', '0'},
			[]byte{'1', '0', '1', '1', '1'},
			[]byte{'1', '1', '1', '1', '1'},
			[]byte{'1', '0', '0', '1', '0'},
		}))

	require.Equal(t, 1,
		maximalSquare1([][]byte{
			[]byte{'0', '1'},
			[]byte{'1', '0'},
		}))
}
