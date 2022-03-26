package lc049

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_groupAnagrams(t *testing.T) {
	require.Equal(t, [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))

	require.Equal(t, [][]string{{"a"}},
		groupAnagrams([]string{"a"}))
}
