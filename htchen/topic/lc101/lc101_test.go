package lc101

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isSymmetric(t *testing.T) {
	require.Equal(t, true, isSymmetric(BuildTree([]int{1, 2, 3, 4, 2, 4, 3}, []int{3, 2, 4, 1, 4, 2, 3})))
	
    require.Equal(t, false, isSymmetric(BuildTree([]int{1, 2, 3, 2, 3}, []int{2,3,1,2,3})))
}
