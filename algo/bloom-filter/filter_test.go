package bloomfilter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_BloomFilter(t *testing.T) {
	b := NewBloomFilter(10, nil, 20)

	b.Add([]byte("abc"))
	require.Equal(t, true, b.Exist([]byte("abc")))

	require.Equal(t, false, b.Exist([]byte("ab")))
}
