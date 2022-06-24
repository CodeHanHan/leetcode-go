package topic75

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sortColors(t *testing.T) {
	var colors []int

	colors = []int{2, 1, 0}
	sortColors(colors)
	require.Equal(t, []int{0, 1, 2}, colors)

	colors = []int{2, 0, 1}
	sortColors(colors)
	require.Equal(t, []int{0, 1, 2}, colors)

	colors = []int{0}
	sortColors(colors)
	require.Equal(t, []int{0}, colors)

	colors = []int{0, 0}
	sortColors(colors)
	require.Equal(t, []int{0, 0}, colors)

	colors = []int{1, 1}
	sortColors(colors)
	require.Equal(t, []int{1, 1}, colors)

	colors = []int{2, 2}
	sortColors(colors)
	require.Equal(t, []int{2, 2}, colors)

	colors = []int{2, 0, 2, 1, 1, 0}
	sortColors(colors)
	require.Equal(t, []int{0, 0, 1, 1, 2, 2}, colors)
}
