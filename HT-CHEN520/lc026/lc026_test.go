package lc26

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeDuplicates(t *testing.T) {

	require.Equal(t, 2, removeDuplicates([]int{1, 1, 2}))

	require.Equal(t, 5, removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
