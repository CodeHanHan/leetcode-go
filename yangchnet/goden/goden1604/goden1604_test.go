package goden1604

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_tictactoe(t *testing.T) {
	require.Equal(t, "X", tictactoe([]string{"O X", " XO", "X O"}))
	require.Equal(t, "Draw", tictactoe([]string{"OOX", "XXO", "OXO"}))
	require.Equal(t, "Pending", tictactoe([]string{"OOX", "XXO", "OX "}))
	require.Equal(t, "X", tictactoe([]string{"X"}))
	require.Equal(t, "Pending", tictactoe([]string{"O X", "X O", "O  "}))
}
