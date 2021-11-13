package lc137

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_singleNumber(t *testing.T) {

	require.Equal(t, 3, singleNumber([]int{2, 2, 3, 2}))

	require.Equal(t, 99, singleNumber([]int{0, 1, 0, 1, 0, 1, 99}))

}

func Test_singleNumber_1(t *testing.T) {

	require.Equal(t, 3, singleNumber([]int{2, 2, 3, 2}))

	require.Equal(t, -4, singleNumber([]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}))

}
