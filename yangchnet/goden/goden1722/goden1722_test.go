package goden1722

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findLadders(t *testing.T) {
	// require.Equal(t,
	// 	[]string{"hit", "hot", "dot", "dog", "cog"},
	// 	findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))

	// require.Equal(t,
	// 	[]string{"hot", "dot", "dog"},
	// 	findLadders("hot", "dog", []string{"hot", "dog", "dot"}))

	// require.Equal(t,
	// 	[]string{"hit", "hot", "dot", "dog", "cog"},
	// 	findLadders("hit", "cog", []string{"hot", "cog", "dot", "dog", "hit", "lot", "log"}))

	require.Equal(t,
		[]string{"hot", "hog", "cog", "dog"},
		findLadders("hot", "dog", []string{"hot", "cog", "dog", "tot", "hog", "hop", "pot", "dot"}))
}
