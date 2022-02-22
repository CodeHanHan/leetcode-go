package topic215

import (
	"container/heap"
	"sort"
)

// 解法1
func findKthLargest(nums []int, k int) int {
	a := sort.IntSlice(nums)
	sort.Sort(sort.Reverse(a))

	return nums[k-1]
}

// 解法2
type IntSlice []int

func (h *IntSlice) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *IntSlice) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntSlice) Len() int {
	return len(*h)
}

func (h *IntSlice) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *IntSlice) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func findKthLargest_2(nums []int, k int) int {
	a := IntSlice(nums)
	heap.Init(&a)

	for i := 1; i < k; i++ {
		heap.Pop(&a)
	}

	return heap.Pop(&a).(int)
}

// 解法3
func findKthLargest_3(nums []int, k int) int {
	return quickSort(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSort(nums []int, lo, hi int, k int) int {
	pivotPos := partition(nums, lo, hi)
	if k == pivotPos {
		return nums[k]
	} else if pivotPos < k {
		return quickSort(nums, pivotPos+1, hi, k)
	}

	return quickSort(nums, lo, pivotPos-1, k)
}

func partition(nums []int, lo, hi int) int {
	pivot := nums[lo]
	for lo < hi {
		for lo < hi && nums[hi] >= pivot {
			hi--
		}
		nums[lo] = nums[hi]
		for lo < hi && nums[lo] <= pivot {
			lo++
		}
		nums[hi] = nums[lo]
	}
	nums[lo] = pivot
	return lo
}
