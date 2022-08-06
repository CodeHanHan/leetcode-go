package topic295

import (
	"container/heap"
	"sort"
)

type MedianFinder struct {
	queMin, queMax hp
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (mf *MedianFinder) AddNum(num int) {
	minQ, maxQ := &mf.queMin, &mf.queMax
	if minQ.Len() == maxQ.Len() {
		heap.Push(minQ, num)

		heap.Push(maxQ, -heap.Pop(minQ).(int))
	} else {
		heap.Push(maxQ, -num)

		heap.Push(minQ, -heap.Pop(maxQ).(int))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	minQ, maxQ := mf.queMin, mf.queMax
	if minQ.Len() == maxQ.Len() {
		return (float64(minQ.Top()) + float64(-maxQ.Top())) / 2
	} else {
		return float64(-maxQ.Top())
	}
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func (h *hp) Top() int {
	return h.IntSlice[0]
}
