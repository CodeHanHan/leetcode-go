package lc131

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_partition(t *testing.T) {
	require.Equal(t, [][]string{{"a", "a", "b"}, {"aa", "b"}},
		partition("aab"))

	require.Equal(t, [][]string{{"a", "a", "b", "a"}, {"a", "aba"}, {"aa", "b", "a"}},
		partition("aaba"))

	require.Equal(t, [][]string{{"a"}},
		partition("a"))
}
