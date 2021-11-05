package lc100

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isSameTree(t *testing.T) {
	require.Equal(t,
		true,
		isSameTree(
			BuildTree([]int{1, 2, 3}, []int{2, 1, 3}),
			BuildTree([]int{1, 2, 3}, []int{2, 1, 3}),
		))

	require.Equal(t,
		false,
		isSameTree(
			BuildTree([]int{1, 2}, []int{2, 1}),
			BuildTree([]int{1, 2}, []int{1, 2}),
		))
}
