package topic27

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeElement(t *testing.T) {
	require.Equal(t, 0, removeElement([]int{}, 1))

	require.Equal(t, 0, removeElement([]int{1}, 1))

	require.Equal(t, 2, removeElement([]int{3, 2, 2, 3}, 3))

	require.Equal(t, 1, removeElement([]int{1, 1, 2}, 1))

	require.Equal(t, 8, removeElement([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 0))
}
