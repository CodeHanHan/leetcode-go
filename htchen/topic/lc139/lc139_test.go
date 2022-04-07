package lc139

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_wordBreak(t *testing.T) {
	require.Equal(t, true, wordBreak("leetcode", []string{"leet", "code"}))

	require.Equal(t, true, wordBreak("applepenapple", []string{"apple", "pen"}))

	require.Equal(t, false, wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}
