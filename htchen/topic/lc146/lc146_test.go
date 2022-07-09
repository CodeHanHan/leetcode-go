package lc146

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Constructor(t *testing.T) {
	lru := Constructor(2)

	lru.Put(1, 1)
	lru.Put(2, 2)
	require.Equal(t, 1, lru.Get(1))

	lru.Put(3, 3)
	require.Equal(t, -1, lru.Get(2))

	lru.Put(4, 4)
	require.Equal(t, -1, lru.Get(1))
	require.Equal(t, 3, lru.Get(3))
	require.Equal(t, 4, lru.Get(4))
}
