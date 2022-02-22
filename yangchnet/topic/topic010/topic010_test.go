package topic010

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isMatch(t *testing.T) {
	require.Equal(t, false, isMatch("aa", "a"))

	require.Equal(t, true, isMatch("aa", "a*"))

	require.Equal(t, true, isMatch("ab", ".*"))

	require.Equal(t, true, isMatch("aab", "c*a*b"))

	require.Equal(t, false, isMatch("mississippi", "mis*is*p*."))
}
