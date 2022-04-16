package goden1716

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_massage(t *testing.T) {
	require.Equal(t, 4, massage([]int{1, 2, 3, 1}))

	require.Equal(t, 12, massage([]int{2, 7, 9, 3, 1}))

	require.Equal(t, 12, massage([]int{2, 1, 4, 5, 3, 1, 1, 3}))
}
