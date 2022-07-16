package topic53

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxSubArray(t *testing.T) {
	require.Equal(t, 6, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
