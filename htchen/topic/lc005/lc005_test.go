package lc005

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestPalindrome1(t *testing.T) {
	require.Equal(t, "bab", longestPalindrome1("babad"))

	require.Equal(t, "bb", longestPalindrome1("cbbd"))
}

func Test_longestPalindrome2(t *testing.T) {
	require.Equal(t, "bab", longestPalindrome1("babad"))

	require.Equal(t, "bb", longestPalindrome1("cbbd"))
}
