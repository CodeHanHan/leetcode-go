package topic028

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_strStr(t *testing.T) {
	require.Equal(t, 2, strStr("hello", "ll"))

	require.Equal(t, -1, strStr("aaaaa", "bba"))

	require.Equal(t, 0, strStr("", ""))

	require.Equal(t, 0, strStr("aaa", "aa"))

	require.Equal(t, 0, strStr("ababa", "aba"))

	require.Equal(t, 2, strStr("aaaba", "aba"))

	require.Equal(t, 0, strStr("a", "a"))
}
