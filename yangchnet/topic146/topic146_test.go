package topic146

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_LRUCache(t *testing.T) {
	lru := Constructor(2)

	lru.Put(1, 1)
	require.Equal(t, 1, lru.Get(1))

	lru.Put(2, 2)
	require.Equal(t, 2, lru.Get(2))
	require.Equal(t, -1, lru.Get(3))

	lru.Put(3, 3)
	require.Equal(t, 3, lru.Get(3))
	require.Equal(t, -1, lru.Get(1))

	lru.Put(3, 4)
	require.Equal(t, 4, lru.Get(3))
}

func Test_LRUCache_1(t *testing.T) {
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
