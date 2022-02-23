package list

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildList(t *testing.T) {

	require.Equal(t, "1 → 2 → 3 → 4 → 5 → nil", BuildListWithHead([]int{1, 2, 3, 4, 5}).String())

	require.Equal(t, []int{1, 2, 3, 4, 5}, BuildListWithNoHead([]int{1, 2, 3, 4, 5}).Slice())

	require.Equal(t, 5, BuildListWithNoHead([]int{1, 2, 3, 4, 5}).Len())

	l1 := BuildListWithNoHead([]int{1, 2, 3, 4})

	require.Equal(t, 4, l1.Tail().Val)

	require.Equal(t, 3, l1.Index(3).Val)
}
