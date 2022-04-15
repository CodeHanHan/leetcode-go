package goden1715

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestWord(t *testing.T) {
	require.Equal(t, "dogwalker", longestWord([]string{"cat", "banana", "dog", "nana", "walk", "walker", "dogwalker"}))
	require.Equal(t, "dogwalker", longestWord([]string{"cat", "banana", "dog", "nana", "catdog", "walk", "walker", "dogwalker"}))
	require.Equal(t, "bananacat", longestWord([]string{"cat", "banana", "bananacat", "dog", "nana", "catdog", "walk", "walker", "dogwalker"}))
	require.Equal(t, "", longestWord([]string{"cat", "banana", "dog", "nana", "walk", "walker"}))
	require.Equal(t, "ccc", longestWord([]string{"ccc", "c"}))
}

func Test_check(t *testing.T) {
	require.Equal(t, true,
		check(map[string]bool{"dog": true, "walker": true, "walk": true}, "dogwalker"))
}
