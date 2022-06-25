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
