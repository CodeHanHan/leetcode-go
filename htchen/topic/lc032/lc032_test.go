package lc032

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestValidParentheses(t *testing.T) {
	require.Equal(t, 2, longestValidParentheses("(()"))
	
    require.Equal(t, 4, longestValidParentheses(")()())"))
	
    require.Equal(t, 0, longestValidParentheses(""))
}
func Test_longestValidParentheses1(t *testing.T) {
	require.Equal(t, 2, longestValidParentheses1("(()"))
	
    require.Equal(t, 4, longestValidParentheses1(")()())"))
	
    require.Equal(t, 0, longestValidParentheses1(""))
}
func Test_longestValidParentheses2(t *testing.T) {
	require.Equal(t, 2, longestValidParentheses2("(()"))
	
    require.Equal(t, 4, longestValidParentheses2(")()())"))
	
    require.Equal(t, 0, longestValidParentheses2(""))
}
