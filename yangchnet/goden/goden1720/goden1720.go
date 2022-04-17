package goden1720

import "container/heap"

type MedianFinder struct {
	maxh *MaxHeap
	minh *MinHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		maxh: new(MaxHeap),
		minh: new(MinHeap),
	}
}

func (this *MedianFinder) AddNum(num int) {
	if len(*this.maxh) > 0 && num <= (*this.maxh)[0] {
		// push to max
		heap.Push(this.maxh, num)
	} else if len(*this.minh) > 0 && (*this.minh)[0] <= num {
		// push to min
		heap.Push(this.minh, num)
	} else {
		heap.Push(this.maxh, num)
	}

	// fix
	for len(*this.maxh)-len(*this.minh) > 1 {
		heap.Push(this.minh, heap.Pop(this.maxh))
	}

	for len(*this.minh)-len(*this.maxh) > 1 {
		heap.Push(this.maxh, heap.Pop(this.minh))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	switch len(*this.minh) - len(*this.maxh) {
	case 0:
		return (float64((*this.maxh)[0]) + float64((*this.minh)[0])) / 2
	case 1:
		return float64((*this.minh)[0])
	case -1:
		return float64((*this.maxh)[0])
	default:
		return (float64((*this.maxh)[0]) + float64((*this.minh)[0])) / 2
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

type MinHeap []int

var _ heap.Interface = (*MinHeap)(nil)

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Less(i int, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *MinHeap) Swap(i int, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Push(x interface{}) {
	(*h) = append((*h), x.(int))
}

func (h *MinHeap) Pop() (v interface{}) {
	(*h), v = (*h)[:len(*h)-1], (*h)[len(*h)-1]
	return
}

type MaxHeap []int

var _ heap.Interface = (*MaxHeap)(nil)

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Less(i int, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) Swap(i int, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(x interface{}) {
	(*h) = append((*h), x.(int))
}

func (h *MaxHeap) Pop() (v interface{}) {
	(*h), v = (*h)[:len(*h)-1], (*h)[len(*h)-1]
	return
}
