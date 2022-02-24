package offer06

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reversePrint(t *testing.T) {
	require.Equal(t, []int{2, 3, 1}, reversePrint(BuildListWithNoHead([]int{1, 3, 2})))

	require.Equal(t, []int{0}, reversePrint(BuildListWithNoHead([]int{0})))
	
	require.Equal(t, []int{2, 2, 5, 1}, reversePrint(BuildListWithNoHead([]int{1, 5, 2, 2})))
}
