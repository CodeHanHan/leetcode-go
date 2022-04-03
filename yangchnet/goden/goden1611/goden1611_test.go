package goden1611

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_divingBoard(t *testing.T) {
	require.Equal(t, []int{3, 4, 5, 6}, divingBoard(1, 2, 3))
	require.Equal(t, []int{}, divingBoard(1, 1, 0))
}

func Test_divingBoard2(t *testing.T) {
	require.Equal(t, []int{3, 4, 5, 6}, divingBoard2(1, 2, 3))
	require.Equal(t, []int{}, divingBoard2(1, 1, 0))
}
