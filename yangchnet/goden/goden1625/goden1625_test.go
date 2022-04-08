package goden1625

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_LRU(t *testing.T) {
	lru := Constructor(2)

	// lru.Put(2, 1)
	// require.Equal(t, 1, lru.Get(2))

	// lru.Put(3, 2)
	// require.Equal(t, -1, lru.Get(2))

	// require.Equal(t, 2, lru.Get(3))

	// require.Equal(t, -1, lru.Get(2))

	lru.Put(2, 1)
	lru.Put(1, 1)
	lru.Put(2, 3)
	lru.Put(4, 1)
	require.Equal(t, -1, lru.Get(1))

	require.Equal(t, 3, lru.Get(2))
}
