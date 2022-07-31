package topic76

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minWindow(t *testing.T) {
	require.Equal(t, "BANC", minWindow("ADOBECODEBANC", "ABC"))

	require.Equal(t, "a", minWindow("a", "a"))

	require.Equal(t, "", minWindow("a", "aa"))

	require.Equal(t, "aba", minWindow("aba", "aa"))

	require.Equal(t, "aa", minWindow("aaa", "aa"))

	require.Equal(t, "", minWindow("aaa", "h"))

	require.Equal(t, "a", minWindow("ab", "a"))

	require.Equal(t, "b", minWindow("ab", "b"))
}

func Test_minWindow1(t *testing.T) {
	require.Equal(t, "BANC", minWindow1("ADOBECODEBANC", "ABC"))

	require.Equal(t, "a", minWindow1("a", "a"))

	require.Equal(t, "", minWindow1("a", "aa"))

	require.Equal(t, "aba", minWindow1("aba", "aa"))

	require.Equal(t, "aa", minWindow1("aaa", "aa"))

	require.Equal(t, "", minWindow1("aaa", "h"))

	require.Equal(t, "a", minWindow1("ab", "a"))

	require.Equal(t, "b", minWindow1("ab", "b"))
}
