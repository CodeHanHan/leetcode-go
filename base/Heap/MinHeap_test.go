package heap

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_MinHeap(t *testing.T) {
	h := new(MinIntHeap)
	for i := 20; i >= 0; i-- {
		heap.Push(h, i) // 向堆中添加元素，使用heap.Push()而不是h.Push()
	}
	heap.Init(h)

	for i := 0; i < 20; i++ {
		require.Equal(t, i, heap.Pop(h).(int)) // 从堆中获取元素，使用heap.Pop()而不是h.Pop()
	}
}
