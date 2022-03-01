package lc061

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_rotateRight(t *testing.T) {
	require.Equal(t, BuildListWithNoHead([]int{4, 5, 1, 2, 3}),
		rotateRight(BuildListWithNoHead([]int{1, 2, 3, 4, 5}), 2))

	require.Equal(t, BuildListWithNoHead([]int{2, 0, 1}),
		rotateRight(BuildListWithNoHead([]int{0, 1, 2}), 4))
}
