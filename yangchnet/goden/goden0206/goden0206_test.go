package goden0206

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPalindrome(t *testing.T) {
	l1 := BuildListWithNoHead([]int{1, 2, 1})
	require.Equal(t, true, isPalindrome(l1))

	l2 := BuildListWithNoHead([]int{1})
	require.Equal(t, true, isPalindrome(l2))

	l3 := BuildListWithNoHead([]int{1, 2, 3, 3, 2, 1})
	require.Equal(t, true, isPalindrome(l3))

}
