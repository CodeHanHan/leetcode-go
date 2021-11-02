package topic148

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sortList(t *testing.T) {
	require.Equal(t,
		BuildListWithNoHead([]int{1, 2, 3, 4}),
		sortList(BuildListWithNoHead([]int{4, 2, 3, 1})))

	require.Equal(t,
		BuildListWithNoHead([]int{}),
		sortList(BuildListWithNoHead([]int{})))

	require.Equal(t,
		BuildListWithNoHead([]int{}),
		sortList(BuildListWithNoHead([]int{})))
}

func Test_sortList_1(t *testing.T) {
	require.Equal(t,
		BuildListWithNoHead([]int{1, 2, 3, 4}),
		sortList_1(BuildListWithNoHead([]int{4, 2, 3, 1})))

	require.Equal(t,
		BuildListWithNoHead([]int{}),
		sortList_1(BuildListWithNoHead([]int{})))

	require.Equal(t,
		BuildListWithNoHead([]int{}),
		sortList_1(BuildListWithNoHead([]int{})))
}

func Test_splitList(t *testing.T) {
	l1, l2 := splitList(BuildListWithNoHead([]int{4, 2, 3, 1}))
	require.Equal(t, "L1: 4 → 2 → nil\n", fmt.Sprintf("L1: %s\n", l1))
	require.Equal(t, "L2: 3 → 1 → nil\n", fmt.Sprintf("L2: %s\n", l2))

	l3, l4 := splitList(BuildListWithNoHead([]int{4, 2, 3, 1, 5}))
	require.Equal(t, "L3: 4 → 2 → 3 → nil\n", fmt.Sprintf("L3: %s\n", l3))
	require.Equal(t, "L4: 1 → 5 → nil\n", fmt.Sprintf("L4: %s\n", l4))

	l5, l6 := splitList(BuildListWithNoHead([]int{4}))
	require.Equal(t, "L5: 4 → nil\n", fmt.Sprintf("L5: %s\n", l5))
	require.Equal(t, "L6: <nil>\n", fmt.Sprintf("L6: %s\n", l6))
}
