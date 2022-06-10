package topic209

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minSubArrayLen(t *testing.T) {
	require.Equal(t, 2, minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))

	require.Equal(t, 1, minSubArrayLen(4, []int{1, 4, 4}))

	require.Equal(t, 0, minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))

	require.Equal(t, 1, minSubArrayLen(4, []int{1, 2, 3, 4}))

	require.Equal(t, 3, minSubArrayLen(11, []int{1, 2, 3, 4, 5}))
}
