package topic022

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_generateParenthesis(t *testing.T) {
	require.Equal(t, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}, generateParenthesis(3))

	require.Equal(t, []string{"()"}, generateParenthesis(1))
}
