package lc207

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canFinish(t *testing.T) {
	require.Equal(t, true, canFinish(2, [][]int{[]int{1, 0}}))
	
    require.Equal(t, false, canFinish(2, [][]int{[]int{1, 0}, []int{0, 1}}))
}
