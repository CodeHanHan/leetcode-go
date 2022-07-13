package topic802

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_eventualSafeNodes(t *testing.T) {
	require.Equal(t, []int{2, 4, 5, 6}, eventualSafeNodes([][]int{
		{1, 2},
		{2, 3},
		{5},
		{0},
		{5},
		{},
		{},
	}))
}
