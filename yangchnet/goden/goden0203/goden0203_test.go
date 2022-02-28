package goden0203

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_deleteNode(t *testing.T) {
	l1 := BuildListWithNoHead([]int{4, 5, 1, 9})
	node := l1.Next
	deleteNode(node)
	require.Equal(t, BuildListWithNoHead([]int{4, 1, 9}), l1)
}
