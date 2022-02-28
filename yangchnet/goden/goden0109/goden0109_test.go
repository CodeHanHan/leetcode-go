package goden0109

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isFlipedString(t *testing.T) {
	require.Equal(t, true, isFlipedString("waterbottle", "erbottlewat"))
	require.Equal(t, true, isFlipedString("aa", "aa"))
	require.Equal(t, true, isFlipedString("aba", "aba"))
	require.Equal(t, true, isFlipedString("", ""))
}
