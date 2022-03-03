package lc086

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_partition(t *testing.T) {
	require.Equal(t, BuildListWithNoHead([]int{1, 2, 2, 4, 3, 5}),
		partition(BuildListWithNoHead([]int{1, 4, 3, 2, 5, 2}), 3))

	require.Equal(t, BuildListWithNoHead([]int{1, 2}),
		partition(BuildListWithNoHead([]int{2, 1}), 2))
}
