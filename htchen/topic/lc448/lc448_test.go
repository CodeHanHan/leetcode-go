package lc448

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findDisappearedNumbers(t *testing.T) {
	require.Equal(t, []int{5, 6}, 
        findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	
    require.Equal(t, []int{2}, 
        findDisappearedNumbers([]int{1, 1}))
}
