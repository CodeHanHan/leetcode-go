package trie

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Trie(t *testing.T) {
	tree := NewTrie()

	tree.Insert("abc")
	tree.Insert("abcc")
	tree.Insert("ab")
	tree.Insert("mnj")
	require.Equal(t, true, tree.Find(("abc")))
	require.Equal(t, true, tree.Find(("ab")))
	require.Equal(t, false, tree.Find(("")))
	require.Equal(t, false, tree.Find(("b")))
	require.Equal(t, false, tree.Find(("mn")))
}
