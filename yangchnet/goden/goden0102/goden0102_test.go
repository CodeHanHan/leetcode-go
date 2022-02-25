package goden0102

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_CheckPermutation(t *testing.T) {
	require.Equal(t, true, CheckPermutation("abc", "bac"))
	require.Equal(t, false, CheckPermutation("abc", "cbac"))
}
