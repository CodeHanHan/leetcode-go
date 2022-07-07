package lc647

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countSubstrings(t *testing.T) {
	require.Equal(t, 3, countSubstrings("abc"))
    
	require.Equal(t, 6, countSubstrings("aaa"))
}
