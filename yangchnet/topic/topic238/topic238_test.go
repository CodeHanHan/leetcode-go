package topic238

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_productExceptSelf(t *testing.T) {
	require.Equal(t, []int{24, 12, 8, 6}, productExceptSelf([]int{1, 2, 3, 4}))

	require.Equal(t, []int{120, 60, 40, 30, 24}, productExceptSelf([]int{1, 2, 3, 4, 5}))
}

func Test_productExceptSelf_1(t *testing.T) {
	require.Equal(t, []int{24, 12, 8, 6}, productExceptSelf_1([]int{1, 2, 3, 4}))

	require.Equal(t, []int{120, 60, 40, 30, 24}, productExceptSelf_1([]int{1, 2, 3, 4, 5}))
}
