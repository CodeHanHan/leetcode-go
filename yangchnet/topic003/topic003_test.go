package topic003

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	require.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))

	require.Equal(t, 1, lengthOfLongestSubstring("aaaaa"))

	require.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))

	require.Equal(t, 3, lengthOfLongestSubstring("dvdf"))

	require.Equal(t, 0, lengthOfLongestSubstring(""))

	require.Equal(t, 3, lengthOfLongestSubstring("abc"))
}
