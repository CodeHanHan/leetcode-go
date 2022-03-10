package backtrack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Bag(t *testing.T) {
	items := []uint32{1, 2, 3, 4, 5, 6, 7, 9}

	b := &Bag{
		Goods:  make([]uint32, 0),
		Cap:    20,
		Weight: 0,
	}
	_, weight := b.Load(items)
	require.Equal(t, uint32(20), weight)
}
