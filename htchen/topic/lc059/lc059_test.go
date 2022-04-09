package lc059

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_generateMatrix(t *testing.T) {
	require.Equal(t, [][]int{[]int{1, 2, 3}, []int{8, 9, 4}, []int{7, 6, 5}},
		generateMatrix(3))
	
        require.Equal(t, [][]int{[]int{1}},
		generateMatrix(1))
}
