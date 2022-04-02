package goden1601

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_swapNumbers(t *testing.T) {
	require.Equal(t, []int{1, 2}, swapNumbers([]int{2, 1}))
}
