package topic8

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_myAtoi(t *testing.T) {
	require.Equal(t, 42, myAtoi("42"))

	require.Equal(t, -42, myAtoi("   -42"))

	require.Equal(t, 42, myAtoi("   +42"))

	require.Equal(t, 0, myAtoi("   "))

	require.Equal(t, 4193, myAtoi("4193 with words"))

	require.Equal(t, 0, myAtoi("words and 987"))

	require.Equal(t, -2147483648, myAtoi("-91283472332"))

	require.Equal(t, 2147483647, myAtoi("91283479999"))

	require.Equal(t, 2147483647, myAtoi("20000000000000000000"))

	require.Equal(t, 12345678, myAtoi("  0000000000012345678"))
}

func Test_myAtoi1(t *testing.T) {
	require.Equal(t, 42, myAtoi1("42"))

	require.Equal(t, -42, myAtoi1("   -42"))

	require.Equal(t, 42, myAtoi1("   +42"))

	require.Equal(t, 0, myAtoi1("   "))

	require.Equal(t, 4193, myAtoi1("4193 with words"))

	require.Equal(t, 0, myAtoi1("words and 987"))

	require.Equal(t, -2147483648, myAtoi1("-91283472332"))

	require.Equal(t, 2147483647, myAtoi1("91283479999"))

	require.Equal(t, 2147483647, myAtoi1("20000000000000000000"))

	require.Equal(t, 12345678, myAtoi1("  0000000000012345678"))
}
