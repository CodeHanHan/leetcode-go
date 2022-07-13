package sort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_quickSort(t *testing.T) {
	var nums []int

	nums = []int{5, 4, 3, 2, 1}
	quickSort(nums)
	require.Equal(t, []int{1, 2, 3, 4, 5}, nums)
}
