package topic45

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_jump(t *testing.T) {
	require.Equal(t, 2, jump([]int{2, 3, 1, 1, 4}))

	require.Equal(t, 2, jump([]int{2, 3, 0, 1, 4}))

	require.Equal(t, 2, jump([]int{3, 1, 1, 1, 4}))

	require.Equal(t, 2, jump([]int{3, 1, 1, 1, 4}))

	require.Equal(t, 2, jump([]int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3}))
}
