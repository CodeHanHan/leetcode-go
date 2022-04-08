package goden1624

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_pairSums(t *testing.T) {
	require.Equal(t, [][]int{
		{5, 6},
	}, pairSums([]int{5, 6, 5}, 11))

	require.Equal(t, [][]int{
		{5, 6},
		{5, 6},
	}, pairSums([]int{5, 6, 5, 6}, 11))
}
