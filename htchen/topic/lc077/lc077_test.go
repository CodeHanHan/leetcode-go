package lc077

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_combine(t *testing.T) {
	require.Equal(t, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}, combine(4, 2))

	require.Equal(t, [][]int{{1}}, combine(1, 1))
}
