package goden0806

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_hanota(t *testing.T) {
	require.Equal(t, []int{3, 2, 1}, hanota([]int{3, 2, 1}, []int{}, []int{}))
}
