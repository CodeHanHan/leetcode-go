package goden0205

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_addTwoNumbers(t *testing.T) {
	require.Equal(t,
		BuildListWithNoHead([]int{2, 1, 9}),
		addTwoNumbers(
			BuildListWithNoHead([]int{7, 1, 6}),
			BuildListWithNoHead([]int{5, 9, 2}),
		),
	)

	require.Equal(t,
		BuildListWithNoHead([]int{3, 5, 2, 5, 5}),
		addTwoNumbers(
			BuildListWithNoHead([]int{1, 2, 3, 4, 5}),
			BuildListWithNoHead([]int{2, 3, 9}),
		),
	)

	require.Equal(t,
		BuildListWithNoHead([]int{0, 1}),
		addTwoNumbers(
			BuildListWithNoHead([]int{5}),
			BuildListWithNoHead([]int{5}),
		),
	)
}
