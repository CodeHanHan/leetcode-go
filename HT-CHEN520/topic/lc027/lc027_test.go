package lc027

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeElement(t *testing.T) {

	require.Equal(t, 2, removeElement([]int{3, 2, 2, 3}, 2))

	require.Equal(t, 5, removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
