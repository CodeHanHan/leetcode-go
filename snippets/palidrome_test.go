package snippets

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPalindrome(t *testing.T) {
	require.Equal(t, true, isPalindrome("a"))
	require.Equal(t, false, isPalindrome("ab"))
	require.Equal(t, true, isPalindrome("aba"))
	require.Equal(t, true, isPalindrome("abaaba"))
}

func Test_expendCenter(t *testing.T) {
	var i, j int

	i, j = expendCenter("a", 0, 0)
	require.Equal(t, 0, i)
	require.Equal(t, 0, j)

	i, j = expendCenter("abccbd", 2, 3)
	require.Equal(t, 1, i)
	require.Equal(t, 4, j)
}
