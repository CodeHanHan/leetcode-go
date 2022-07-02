package lc394

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_decodeString(t *testing.T) {
	require.Equal(t, "aaabcbc", decodeString("3[a]2[bc]"))
	
    require.Equal(t, "accaccacc", decodeString("3[a2[c]]"))
	
    require.Equal(t, "abcabccdcdcdef", decodeString("2[abc]3[cd]ef"))
	
    require.Equal(t, "abccdcdcdxyz", decodeString("abc3[cd]xyz"))
}
