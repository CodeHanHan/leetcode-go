package goden0106

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_compressString(t *testing.T) {
	require.Equal(t, "a2b1c5a3", compressString("aabcccccaaa"))
	require.Equal(t, "abbccd", compressString("abbccd"))
	require.Equal(t, "a", compressString("a"))
	require.Equal(t, "", compressString(""))
	require.Equal(t, "abc", compressString("abc"))
	require.Equal(t, "aa", compressString("aa"))
}
