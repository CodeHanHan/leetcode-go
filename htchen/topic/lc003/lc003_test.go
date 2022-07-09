package lc003

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	require.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))

	require.Equal(t, 1, lengthOfLongestSubstring("bbbb"))

	require.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
}
