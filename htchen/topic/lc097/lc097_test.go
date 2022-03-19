package lc097

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isInterleave(t *testing.T) {
	require.Equal(t, true, isInterleave("aabcc", "dbbca", "aadbbcbcac"))

	require.Equal(t, false, isInterleave("aabcc", "dbbca", "aadbbbaccc"))

	require.Equal(t, true, isInterleave("", "", ""))
}
