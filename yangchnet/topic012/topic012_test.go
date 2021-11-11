package topic012

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_intToRoman(t *testing.T) {
	require.Equal(t, "XC", intToRoman(90))

	require.Equal(t, "III", intToRoman(3))

	require.Equal(t, "IV", intToRoman(4))

	require.Equal(t, "M", intToRoman(1000))

	require.Equal(t, "IX", intToRoman(9))

	require.Equal(t, "MCCXXXIV", intToRoman(1234))

	require.Equal(t, "MCCXCIX", intToRoman(1299))

	require.Equal(t, "XL", intToRoman(40))

	require.Equal(t, "XC", intToRoman(90))

	require.Equal(t, "CM", intToRoman(900))

	require.Equal(t, "CD", intToRoman(400))

	require.Equal(t, "XLIV", intToRoman(44))

	require.Equal(t, "XCIX", intToRoman(99))

	require.Equal(t, "CMXCIX", intToRoman(999))
}
