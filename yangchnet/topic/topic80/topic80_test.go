package topic80

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeDuplicates(t *testing.T) {
	var nums []int

	nums = []int{1, 1, 1, 2, 2, 3}
	require.Equal(t, 5, removeDuplicates(nums))
	require.Equal(t, []int{1, 1, 2, 2, 3}, nums[:removeDuplicates(nums)-1])
}
