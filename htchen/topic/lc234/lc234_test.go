package lc234

import (
	"testing"

	list "github.com/CodeHanHan/leetcode-go/base/LinkList"
	"github.com/stretchr/testify/require"
)

func Test_isPalindrome1(t *testing.T) {
	require.Equal(t, true,
		isPalindrome1(list.BuildListWithNoHead([]int{1, 2, 2, 1})))

	require.Equal(t, false,
		isPalindrome1(list.BuildListWithNoHead([]int{1, 2})))
}

func Test_isPalindrome2(t *testing.T) {
	require.Equal(t, true,
		isPalindrome2(list.BuildListWithNoHead([]int{1, 2, 2, 1})))

	require.Equal(t, false,
		isPalindrome2(list.BuildListWithNoHead([]int{1, 2})))
}
