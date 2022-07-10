package topic120

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumTotal(t *testing.T) {
	require.Equal(t, 11, minimumTotal([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}))

	require.Equal(t, -10, minimumTotal([][]int{
		{-10},
	}))
}

func Test_minimumTotal1(t *testing.T) {
	require.Equal(t, 11, minimumTotal1([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}))

	require.Equal(t, -10, minimumTotal1([][]int{
		{-10},
	}))
}
