package goden0202

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_kthToLast(t *testing.T) {
	require.Equal(t, 4, kthToLast(BuildListWithNoHead([]int{1, 2, 3, 4, 5}), 2))
	require.Equal(t, 5, kthToLast(BuildListWithNoHead([]int{1, 2, 3, 4, 5}), 1))
}
