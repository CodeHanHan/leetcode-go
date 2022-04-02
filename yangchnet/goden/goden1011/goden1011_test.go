package goden1011

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_wiggleSort(t *testing.T) {
	var arr []int

	arr = []int{5, 3, 1, 2, 3}
	wiggleSort(arr)
	require.Equal(t, []int{2, 1, 3, 3, 5}, arr)

	arr = []int{}
	wiggleSort(arr)
	require.Equal(t, []int{}, arr)
}
