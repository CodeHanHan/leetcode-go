package topic2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_addTwoNumbers(t *testing.T) {
	l1 := BuildListWithNoHead([]int{2, 4, 3})
	l2 := BuildListWithNoHead([]int{5, 6, 4})

	l := addTwoNumbers_1(l1, l2)
	require.Equal(t, "[7 0 8]", fmt.Sprintf("%v", l.Slice()))

	l1_1 := BuildListWithNoHead([]int{5, 6})
	l2_1 := BuildListWithNoHead([]int{1, 2, 3})
	l_1 := addTwoNumbers_1(l1_1, l2_1)
	require.Equal(t, "[6 8 3]", fmt.Sprintf("%v", l_1.Slice()))

	l1_2 := BuildListWithNoHead([]int{9, 9, 9, 9, 9, 9, 9})
	l2_2 := BuildListWithNoHead([]int{9, 9, 9, 9})
	l_2 := addTwoNumbers_1(l1_2, l2_2)
	require.Equal(t, "[8 9 9 9 0 0 0 1]", fmt.Sprintf("%v", l_2.Slice()))

	l1_3 := BuildListWithNoHead([]int{0})
	l2_3 := BuildListWithNoHead([]int{9, 9, 9, 9})
	l_3 := addTwoNumbers_1(l1_3, l2_3)
	require.Equal(t, "[9 9 9 9]", fmt.Sprintf("%v", l_3.Slice()))
}
