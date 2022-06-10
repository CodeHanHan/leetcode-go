package topic74

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_searchMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60},
	}

	require.Equal(t, false, searchMatrix(matrix, 13))
	require.Equal(t, true, searchMatrix(matrix, 23))
}
