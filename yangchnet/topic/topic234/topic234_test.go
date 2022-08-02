package topic234

import (
	"testing"

	. "github.com/CodeHanHan/leetcode-go/base/LinkList"
	"github.com/stretchr/testify/require"
)

func Test_isPalindrome(t *testing.T) {
	var list *ListNode

	list = BuildListWithNoHead([]int{1, 2, 1})
	require.Equal(t, true, isPalindrome(list))

	list = BuildListWithNoHead([]int{1, 2})
	require.Equal(t, false, isPalindrome(list))

	list = BuildListWithNoHead([]int{1})
	require.Equal(t, true, isPalindrome(list))

	list = BuildListWithNoHead([]int{})
	require.Equal(t, true, isPalindrome(list))
}
