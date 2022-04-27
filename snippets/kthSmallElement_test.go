package snippets

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_kthSmallElement(t *testing.T) {
	require.Equal(t,
		2,
		kthSmallElement([]int{1, 2}, []int{3, 4}, 2),
	)

	require.Equal(t,
		3,
		kthSmallElement([]int{1, 2}, []int{3, 4}, 3),
	)

	require.Equal(t,
		1,
		kthSmallElement([]int{1, 2}, []int{3, 4}, 1),
	)

	require.Equal(t,
		byte('a'),
		kthSmallElement([]byte{'a', 'b'}, []byte{'c', 'd'}, 1),
	)
}
