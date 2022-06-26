package topic84

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_largestRectangleArea(t *testing.T) {
	require.Equal(t, 10, largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))

	require.Equal(t, 2, largestRectangleArea([]int{1, 1}))
}
