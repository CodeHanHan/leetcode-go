package goden1622

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_printKMoves(t *testing.T) {
	require.Equal(t, []string{
		"_X",
		"LX",
	}, printKMoves(2))

	require.Equal(t, []string{
		"_U",
		"X_",
		"XX",
	}, printKMoves(5))

	require.Equal(t, []string{
		"R",
	}, printKMoves(0))
}
