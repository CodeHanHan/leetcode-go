package topic217

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_containsDuplicate(t *testing.T) {
	require.Equal(t, true, containsDuplicate([]int{1, 1}))

	require.Equal(t, true, containsDuplicate([]int{1, 1, 2}))

	require.Equal(t, false, containsDuplicate([]int{1, 3, 2}))

	require.Equal(t, false, containsDuplicate([]int{1}))

	require.Equal(t, false, containsDuplicate([]int{}))
}
