package topic151

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverseWords(t *testing.T) {
	require.Equal(t, "blue is sky the", reverseWords("the sky is blue"))

	require.Equal(t, "world hello", reverseWords("  hello world  "))

	require.Equal(t, "example good a", reverseWords("a good   example"))

	require.Equal(t, "a", reverseWords("a"))
}

func Test_reverseWords1(t *testing.T) {
	require.Equal(t, "blue is sky the", reverseWords1("the sky is blue"))

	require.Equal(t, "world hello", reverseWords1("  hello world  "))

	require.Equal(t, "example good a", reverseWords1("a good   example"))

	require.Equal(t, "a", reverseWords1("a"))
}
