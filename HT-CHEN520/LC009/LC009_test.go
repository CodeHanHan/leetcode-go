package LC009

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPalindrome(t *testing.T) {
	require.Equal(t, true, isPalindrome(121))

	require.Equal(t, false, isPalindrome(-121))

	require.Equal(t, false, isPalindrome(10))

	require.Equal(t, false, isPalindrome(-101))
}
