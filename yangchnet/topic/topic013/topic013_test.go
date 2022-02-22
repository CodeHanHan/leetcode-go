package topic013

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_romanToInt(t *testing.T) {
	require.Equal(t, 3, romanToInt("III"))

	require.Equal(t, 4, romanToInt("IV"))

	require.Equal(t, 9, romanToInt("IX"))

	require.Equal(t, 58, romanToInt("LVIII"))

	require.Equal(t, 1994, romanToInt("MCMXCIV"))

	require.Equal(t, 99, romanToInt("XCIX"))

	require.Equal(t, 900, romanToInt("CM"))
}

func Test_romanToInt_1(t *testing.T) {
	require.Equal(t, 3, romanToInt_1("III"))

	require.Equal(t, 4, romanToInt_1("IV"))

	require.Equal(t, 9, romanToInt_1("IX"))

	require.Equal(t, 58, romanToInt_1("LVIII"))

	require.Equal(t, 1994, romanToInt_1("MCMXCIV"))

	require.Equal(t, 99, romanToInt_1("XCIX"))

	require.Equal(t, 900, romanToInt_1("CM"))
}
