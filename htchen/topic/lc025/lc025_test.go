package lc025

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverseKGroup(t *testing.T) {
	require.Equal(t, BuildListWithNoHead([]int{2, 1, 4, 3, 5}),
		reverseKGroup(BuildListWithNoHead([]int{1, 2, 3, 4, 5}), 2))

	require.Equal(t, BuildListWithNoHead([]int{3, 2, 1, 4, 5}),
		reverseKGroup(BuildListWithNoHead([]int{1, 2, 3, 4, 5}), 3))

	require.Equal(t, BuildListWithNoHead([]int{1, 2, 3, 4, 5}),
		reverseKGroup(BuildListWithNoHead([]int{1, 2, 3, 4, 5}), 1))
}
