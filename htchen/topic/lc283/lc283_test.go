package lc283

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_moveZeroes(t *testing.T) {
	a := []int{0, 1, 0, 3, 12}
	moveZeroes(a)
	require.Equal(t, []int{1, 3, 12, 0, 0}, a)

	b := []int{0}
	moveZeroes(b)
	require.Equal(t, []int{0}, b)
}
