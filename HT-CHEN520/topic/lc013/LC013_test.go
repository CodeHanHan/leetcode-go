package LC013

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
}
