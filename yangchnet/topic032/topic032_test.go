package topic032

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestValidParentheses(t *testing.T) {
	require.Equal(t, 2, longestValidParentheses("(()"))

	require.Equal(t, 4, longestValidParentheses(")()())"))

	require.Equal(t, 0, longestValidParentheses(""))

	require.Equal(t, 6, longestValidParentheses(")()(())("))

	require.Equal(t, 2, longestValidParentheses(")())())"))
}
