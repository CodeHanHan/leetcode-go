package lc083

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_deleteDuplicates(t *testing.T) {
	require.Equal(t, BuildListWithNoHead([]int{1, 2}),
		deleteDuplicates(BuildListWithNoHead([]int{1, 1, 2})))

	require.Equal(t, BuildListWithNoHead([]int{1, 2, 3}),
		deleteDuplicates(BuildListWithNoHead([]int{1, 1, 2, 3, 3})))
}
