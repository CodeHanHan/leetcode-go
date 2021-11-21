package lc075

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sortColors(t *testing.T) {
	a := []int{2, 0, 2, 1, 1, 0}
	sortColors(a)
	require.Equal(t, []int{0, 0, 1, 1, 2, 2}, a)

	b := []int{2, 0, 1}
	sortColors(b)
	require.Equal(t, []int{0, 1, 2}, a)
}
