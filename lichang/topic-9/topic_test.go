package topic9

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPalindrome(t *testing.T) {
	require.Equal(t, true, isPalindrome(121))

	require.Equal(t, false, isPalindrome(123))

	require.Equal(t, false, isPalindrome(10))

	require.Equal(t, false, isPalindrome(-121))

	require.Equal(t, true, isPalindrome(1))

	require.Equal(t, true, isPalindrome(0))
}

func Test_isPalindrome_1(t *testing.T) {
	require.Equal(t, true, isPalindrome_1(121))

	require.Equal(t, false, isPalindrome_1(123))

	require.Equal(t, false, isPalindrome_1(10))

	require.Equal(t, false, isPalindrome_1(-121))

	require.Equal(t, true, isPalindrome_1(1))

	require.Equal(t, true, isPalindrome_1(9))

	require.Equal(t, true, isPalindrome_1(0))
}
