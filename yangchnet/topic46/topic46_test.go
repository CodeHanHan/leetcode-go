package topic46

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_permute(t *testing.T) {
	require.Equal(t, [][]int{{1, 2}, {2, 1}}, permute([]int{1, 2}))

	require.Equal(t, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}, permute([]int{1, 2, 3}))

	require.Equal(t, [][]int{{1}}, permute([]int{1}))
}
