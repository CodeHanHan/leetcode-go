package topic035

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_searchInsert(t *testing.T) {
	require.Equal(t, 2, searchInsert([]int{1, 3, 5, 6}, 5))

	require.Equal(t, 1, searchInsert([]int{1, 3, 5, 6}, 2))

	require.Equal(t, 4, searchInsert([]int{1, 3, 5, 6}, 7))

	require.Equal(t, 0, searchInsert([]int{1, 3, 5, 6}, 0))

	require.Equal(t, 0, searchInsert([]int{1}, 0))
}

func Test_searchInsert_1(t *testing.T) {
	require.Equal(t, 2, searchInsert_1([]int{1, 3, 5, 6}, 5))

	require.Equal(t, 1, searchInsert_1([]int{1, 3, 5, 6}, 2))

	require.Equal(t, 4, searchInsert_1([]int{1, 3, 5, 6}, 7))

	require.Equal(t, 0, searchInsert_1([]int{1, 3, 5, 6}, 0))

	require.Equal(t, 0, searchInsert_1([]int{1}, 0))
}
