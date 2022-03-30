package goden1001

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_merge(t *testing.T) {
	var a, b []int

	a = []int{1, 2, 3, 0, 0, 0}
	b = []int{2, 5, 6}
	merge(a, 3, b, 3)
	require.Equal(t, []int{1, 2, 2, 3, 5, 6}, a)

	a = []int{0}
	b = []int{1}
	merge(a, 0, b, 1)
	require.Equal(t, []int{1}, a)
}
