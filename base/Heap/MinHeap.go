package heap

import "container/heap"

type MinIntHeap []int

var _ heap.Interface = (*MinIntHeap)(nil)

func (h *MinIntHeap) Len() int {
	return len(*h)
}

func (h *MinIntHeap) Less(i int, j int) bool {
	return (*h)[i] < (*h)[j]
}

// Swap swaps the elements with indexes i and j.
func (h *MinIntHeap) Swap(i int, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinIntHeap) Push(x interface{}) {
	(*h) = append((*h), x.(int)) // 最小的值本来是存储在index=0位置的，但在pop之前将其与最后一个元素交换了位置
}

func (h *MinIntHeap) Pop() (v interface{}) {
	(*h), v = (*h)[:len(*h)-1], (*h)[len(*h)-1]
	return
}
