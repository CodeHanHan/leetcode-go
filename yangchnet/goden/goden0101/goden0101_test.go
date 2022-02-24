package goden0101

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isUnique(t *testing.T) {
	require.Equal(t, true, isUnique("abcdefg"))
	require.Equal(t, false, isUnique("abcdaefg"))
	require.Equal(t, false, isUnique("leetcode"))
	require.Equal(t, false, isUnique("ee"))
}
