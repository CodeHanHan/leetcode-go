package topic018

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_fourSum(t *testing.T) {
	require.Equal(t, [][]int{{2, 2, 2, 2}}, fourSum([]int{2, 2, 2, 2, 2}, 8))

	require.Equal(t, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}, fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}
