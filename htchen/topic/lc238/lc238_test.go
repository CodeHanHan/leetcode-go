package lc238

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_productExceptSelf(t *testing.T) {
	require.Equal(t, []int{24, 12, 8, 6},
		productExceptSelf1([]int{1, 2, 3, 4}))

	require.Equal(t, []int{0, 0, 9, 0, 0},
		productExceptSelf1([]int{-1, 1, 0, -3, 3}))
}

func Test_productExceptSelf2(t *testing.T) {
	require.Equal(t, []int{24, 12, 8, 6},
		productExceptSelf2([]int{1, 2, 3, 4}))

	require.Equal(t, []int{0, 0, 9, 0, 0},
		productExceptSelf2([]int{-1, 1, 0, -3, 3}))
}
