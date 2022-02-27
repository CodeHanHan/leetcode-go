package goden0104

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canPermutePalindrome(t *testing.T) {
	require.Equal(t, true, canPermutePalindrome("tactcoa"))
	require.Equal(t, true, canPermutePalindrome("aabbcc"))
	require.Equal(t, true, canPermutePalindrome("aabbccd"))
	require.Equal(t, true, canPermutePalindrome("aabbccddd"))

	require.Equal(t, false, canPermutePalindrome("aabbccdddfff"))

	require.Equal(t, false, canPermutePalindrome("ab"))
	require.Equal(t, false, canPermutePalindrome("abbb"))
	require.Equal(t, false, canPermutePalindrome("abc"))
	require.Equal(t, false, canPermutePalindrome("AaBb//a"))
}
