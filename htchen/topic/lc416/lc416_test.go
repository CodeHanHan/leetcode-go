package lc416

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canPartition(t *testing.T) {
	require.Equal(t, true, canPartition([]int{1, 5, 11, 5}))
	
	require.Equal(t, false, canPartition([]int{1, 2, 3, 5}))
}
