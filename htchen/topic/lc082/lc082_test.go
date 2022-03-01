package lc082

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_deleteDuplicates(t *testing.T) {
	require.Equal(t, BuildListWithNoHead([]int{1, 2, 5}),
		deleteDuplicates(BuildListWithNoHead([]int{1, 2, 3, 3, 4, 4, 5})))

	require.Equal(t, BuildListWithNoHead([]int{2, 3}),
		deleteDuplicates(BuildListWithNoHead([]int{1, 1, 1, 2, 3})))
}
