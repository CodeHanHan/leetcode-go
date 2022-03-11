package dynamic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_MinDist(t *testing.T) {
	matrix := [][]int{
		{1, 3, 5, 9},
		{2, 1, 3, 4},
		{5, 2, 6, 7},
		{6, 8, 4, 3},
	}

	require.Equal(t, 19, MinDist(matrix))
}
