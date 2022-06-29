package topic97

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isInterleave(t *testing.T) {
	require.Equal(t, true, isInterleave(
		"aabcc",
		"dbbca",
		"aadbbcbcac",
	))

	require.Equal(t, false, isInterleave(
		"aabcc",
		"dbbca",
		"aadbbbaccc",
	))

	require.Equal(t, true, isInterleave(
		"aabcc",
		"dbbca",
		"aadbcbbcac",
	))

	require.Equal(t, true, isInterleave(
		"",
		"",
		"",
	))

	require.Equal(t, false, isInterleave(
		"",
		"",
		"a",
	))

	require.Equal(t, false, isInterleave(
		"",
		"abc",
		"ab",
	))
}
