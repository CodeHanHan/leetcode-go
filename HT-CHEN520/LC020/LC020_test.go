package LC020

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isValid(t *testing.T) {
	require.Equal(t, true, isValid("()"))

	require.Equal(t, true, isValid("()[]{}"))

	require.Equal(t, false, isValid("(]"))

	require.Equal(t, false, isValid("([)]"))

	require.Equal(t, true, isValid("{[]}"))
}
