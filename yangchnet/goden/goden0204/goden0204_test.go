package goden0204

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_partition(t *testing.T) {
	l1 := BuildListWithNoHead([]int{1, 4, 3, 2, 5, 2})
	require.Equal(t,
		BuildListWithNoHead([]int{1, 2, 2, 4, 3, 5}),
		partition(l1, 3),
	)

	l2 := BuildListWithNoHead([]int{2, 1})
	require.Equal(t, BuildListWithNoHead([]int{1, 2}), partition(l2, 2))
}
