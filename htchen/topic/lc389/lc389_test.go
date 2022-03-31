package lc389

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findTheDifference1(t *testing.T) {
	require.Equal(t, byte('e'), findTheDifference1("abcd", "abcde"))

	require.Equal(t, byte('y'), findTheDifference1("", "y"))
}

func Test_findTheDifference2(t *testing.T) {
	require.Equal(t, byte('e'), findTheDifference2("abcd", "abcde"))

	require.Equal(t, byte('y'), findTheDifference2("", "y"))
}
