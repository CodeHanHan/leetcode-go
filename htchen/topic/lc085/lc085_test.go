package lc085

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximalRectangle(t *testing.T) {
	require.Equal(t, 6,
		maximalRectangle([][]byte{
			[]byte{'1', '0', '1', '0', '0'},
			[]byte{'1', '0', '1', '1', '1'},
			[]byte{'1', '1', '1', '1', '1'},
			[]byte{'1', '0', '0', '1', '0'}}))
}
