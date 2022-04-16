package goden1717

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_multiSearch(t *testing.T) {
	require.Equal(t, [][]int{{1, 4}, {8}, {}, {3}, {1, 4, 7, 10}, {5}}, multiSearch(
		"mississippi",
		[]string{"is", "ppi", "hi", "sis", "i", "ssippi"},
	))

	require.Equal(t, [][]int{{}, {}, {}, {}, {}, {}}, multiSearch(
		"",
		[]string{"is", "ppi", "hi", "sis", "i", "ssippi"},
	))

	require.Equal(t, [][]int{{}}, multiSearch(
		"abc",
		[]string{""},
	))
}
