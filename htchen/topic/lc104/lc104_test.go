package lc104

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxDepth(t *testing.T) {
	require.Equal(t, 3, maxDepth(BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})))
}
