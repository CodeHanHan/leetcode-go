package topic55

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canJump(t *testing.T) {
	require.Equal(t, true, canJump([]int{2, 3, 1, 1, 4}))

	require.Equal(t, false, canJump([]int{3, 2, 1, 0, 4}))

	require.Equal(t, true, canJump([]int{3, 2, 1, 1, 4}))
}
