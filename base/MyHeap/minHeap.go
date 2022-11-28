package heap

import "golang.org/x/exp/constraints"

type MinHeap[T constraints.Ordered] struct {
	vars   []T // 数据存储位置
	length int
}

func NewMinHeap[T constraints.Ordered](vars []T, n int) *MinHeap[T] {
	var noop T
	h := MinHeap[T]{
		vars:   append([]T{noop}, vars...),
		length: n,
	}

	for i := n / 2; i >= 1; i-- {
		h.heapify(h.length, i)
	}

	return &h
}

func (h *MinHeap[T]) Insert(var_ T) {
	h.vars = append(h.vars[:1], append([]T{var_}, h.vars[1:]...)...)
	h.length += 1

	h.heapify(h.length, 1)
}

func (h *MinHeap[T]) Pop() T {
	h.vars[1], h.vars[h.length] = h.vars[h.length], h.vars[1]

	var_ := h.vars[h.length]

	h.vars = h.vars[:h.length]
	h.length -= 1

	h.heapify(h.length, 1)

	return var_
}

func (h *MinHeap[T]) Sorted() {
	k := h.length
	for k > 1 {
		h.vars[1], h.vars[k] = h.vars[k], h.vars[1]
		k--
		h.heapify(k, 1)
	}
}

func (h *MinHeap[T]) heapify(n, i int) {
	for {
		minPos := i
		if i*2 <= n && h.vars[i*2] < h.vars[minPos] {
			minPos = i * 2
		}

		if i*2+1 <= n && h.vars[i*2+1] < h.vars[minPos] {
			minPos = i*2 + 1
		}

		if minPos == i {
			break
		}

		h.vars[i], h.vars[minPos] = h.vars[minPos], h.vars[i]
		i = minPos
	}
}
