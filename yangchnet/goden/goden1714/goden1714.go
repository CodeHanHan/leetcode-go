package goden1714

import (
	"container/heap"
	"sort"
)

type myHeap []int

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Pop() (v interface{}) {
	*h, v = (*h)[:len(*h)-1], (*h)[len(*h)-1]
	return
}

func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func smallestK(arr []int, k int) []int {
	h := new(myHeap)
	for _, a := range arr {
		heap.Push(h, a)
	}

	ret := make([]int, k)
	for i := range ret {
		ret[i] = heap.Pop(h).(int)
	}

	return ret
}

// 排序法
func smallestK2(arr []int, k int) []int {
	sort.Ints(arr)
	ret := make([]int, k)
	copy(ret, arr)

	return ret
}
