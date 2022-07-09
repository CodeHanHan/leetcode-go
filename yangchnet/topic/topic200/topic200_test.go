package topic200

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numIslands(t *testing.T) {
	var land [][]byte

	land = [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	require.Equal(t, 1, numIslands(land))

	land = [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	require.Equal(t, 3, numIslands(land))
}
