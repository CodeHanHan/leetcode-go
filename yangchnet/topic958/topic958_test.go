package topic958

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isCompleteTree(t *testing.T) {
	require.Equal(t, true, isCompleteTree(BuildTree([]int{1, 2, 4, 5, 3, 6}, []int{4, 2, 5, 1, 6, 3})))

	require.Equal(t, false, isCompleteTree(BuildTree([]int{1, 2, 4, 5, 3, 7}, []int{4, 2, 5, 1, 3, 7})))

	require.Equal(t, false, isCompleteTree(BuildTree([]int{1, 2, 3, 4}, []int{2, 1, 4, 3})))
}
