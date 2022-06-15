package topic287

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findDuplicate(t *testing.T) {
	require.Equal(t, 2, findDuplicate([]int{1, 3, 4, 2, 2}))
	require.Equal(t, 3, findDuplicate([]int{3, 1, 3, 4, 2}))
	require.Equal(t, 3, findDuplicate([]int{3, 3, 3, 3, 3}))
}
