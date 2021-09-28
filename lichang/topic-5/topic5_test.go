package topic5

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPalindrome(t *testing.T) {
	require.Equal(t, true, isPalindrome("a"))

	require.Equal(t, false, isPalindrome("ab"))

	require.Equal(t, false, isPalindrome("abc"))

	require.Equal(t, true, isPalindrome("aba"))

	require.Equal(t, true, isPalindrome("bbbb"))

	require.Equal(t, false, isPalindrome("bbbba"))
}

func Test_longestPalindrome(t *testing.T) {
	require.Equal(t, "bab", longestPalindrome("babad"))

	require.Equal(t, "bb", longestPalindrome("cbbd"))

	require.Equal(t, "a", longestPalindrome("a"))

	require.Equal(t, "a", longestPalindrome("ac"))

	require.Equal(t, "cbc", longestPalindrome("acbc"))

	require.Equal(t, "abcccba", longestPalindrome("abcccba"))

	require.Equal(t, "ccc", longestPalindrome("ccc"))
}

func Test_longestPalindrome_1(t *testing.T) {
	require.Equal(t, "bab", longestPalindrome_1("babad"))

	require.Equal(t, "bb", longestPalindrome_1("cbbd"))

	require.Equal(t, "a", longestPalindrome_1("a"))

	require.Equal(t, "a", longestPalindrome_1("ac"))

	require.Equal(t, "cbc", longestPalindrome_1("acbc"))

	require.Equal(t, "abcccba", longestPalindrome_1("abcccba"))

	require.Equal(t, "ccc", longestPalindrome_1("ccc"))
}
