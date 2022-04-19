package lc200

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numIslands(t *testing.T) {
	require.Equal(t, 1,
		numIslands(
			[][]byte{
				[]byte{'1', '1', '1', '1', '0'},
				[]byte{'1', '1', '0', '1', '0'},
				[]byte{'1', '1', '0', '0', '0'},
				[]byte{'0', '0', '0', '0', '0'},
			}))

	require.Equal(t, 3,
		numIslands(
			[][]byte{
				[]byte{'1', '1', '0', '0', '0'},
				[]byte{'1', '1', '0', '0', '0'},
				[]byte{'0', '0', '1', '0', '0'},
				[]byte{'0', '0', '0', '1', '1'},
			}))
}
