package lc242

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isAnagram1(t *testing.T) {
	require.Equal(t, true, isAnagram1("anagram", "nagaram"))
	require.Equal(t, false, isAnagram1("rat", "car"))
}

func Test_isAnagram2(t *testing.T) {
	require.Equal(t, true, isAnagram2("anagram", "nagaram"))
	require.Equal(t, false, isAnagram2("rat", "car"))
}