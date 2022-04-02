package bit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_bit(t *testing.T) {
	bit := NewBinaryIndexedTree([]int{1, 2, 2, 4, 5})
	require.Equal(t, 5, bit.PrefixSum(2))
	bit.Update(2, 1)
	require.Equal(t, 6, bit.PrefixSum(2))
	require.Equal(t, 10, bit.PrefixSum(3))
	require.Equal(t, 9, bit.RangeSum(1, 3))
}
