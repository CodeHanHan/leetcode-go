package topic119

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getRow(t *testing.T) {
	require.Equal(t, []int{1, 4, 6, 4, 1}, getRow(4))

	require.Equal(t, []int{1}, getRow(0))

	require.Equal(t, []int{1, 3, 3, 1}, getRow(3))

	require.Equal(t, []int{1, 2, 1}, getRow(2))

	require.Equal(t, []int{1, 1}, getRow(1))
}
