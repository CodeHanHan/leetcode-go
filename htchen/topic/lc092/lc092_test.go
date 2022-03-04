package lc092

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverseBetween1(t *testing.T) {
	require.Equal(t, BuildListWithNoHead([]int{1, 4, 3, 2, 5}),
		reverseBetween1(BuildListWithNoHead([]int{1, 2, 3, 4, 5}), 2, 4))

	require.Equal(t, BuildListWithNoHead([]int{5}),
		reverseBetween1(BuildListWithNoHead([]int{5}), 1, 1))
}

func Test_reverseBetween2(t *testing.T) {
	require.Equal(t, BuildListWithNoHead([]int{1, 4, 3, 2, 5}),
		reverseBetween2(BuildListWithNoHead([]int{1, 2, 3, 4, 5}), 2, 4))

	require.Equal(t, BuildListWithNoHead([]int{5}),
		reverseBetween2(BuildListWithNoHead([]int{5}), 1, 1))
}
