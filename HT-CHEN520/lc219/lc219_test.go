package lc219

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_containsNearbyDuplicate(t *testing.T) {

	require.Equal(t, true, containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))

	require.Equal(t, true, containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))

	require.Equal(t, false, containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
