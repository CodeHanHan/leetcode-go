package lc206

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverseList(t *testing.T) {
	require.Equal(t, BuildListWithNoHead([]int{5, 4, 3, 2, 1}),
		reverseList(BuildListWithNoHead([]int{1, 2, 3, 4, 5})))

	require.Equal(t, BuildListWithNoHead([]int{2, 1}),
		reverseList(BuildListWithNoHead([]int{1, 2})))

	require.Equal(t, BuildListWithNoHead([]int{}),
		reverseList(BuildListWithNoHead([]int{})))
}
