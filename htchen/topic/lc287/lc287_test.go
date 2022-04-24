package lc287

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findDuplicate1(t *testing.T) {
	require.Equal(t, 2, findDuplicate1([]int{1, 3, 4, 2, 2}))
	
    require.Equal(t, 3, findDuplicate1([]int{3, 1, 3, 4, 2}))
}

func Test_findDuplicate2(t *testing.T) {
	require.Equal(t, 2, findDuplicate2([]int{1, 3, 4, 2, 2}))
	
    require.Equal(t, 3, findDuplicate2([]int{3, 1, 3, 4, 2}))
}