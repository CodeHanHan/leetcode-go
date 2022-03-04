package lc215

import "sort"

func findKthLargest1(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	return nums[n-k]
}

func findKthLargest2(nums []int, k int) int {
	n := len(nums)
	return quickSort(nums, 0, n-1, n-k)
}

func quickSort(nums []int, low, high int, k int) int {
	pivotindex := partition(nums, low, high)
	if pivotindex == k {
		return nums[k]
	} else if pivotindex < k {
		return quickSort(nums, pivotindex+1, high, k)
	} else {
		return quickSort(nums, low, pivotindex-1, k)
	}
}

func partition(nums []int, low int, high int) int {
	key := nums[low]
	for low < high {
		for low < high && nums[high] >= key {
			high--
		}
		nums[low] = nums[high]

		for low < high && nums[low] <= key {
			low++
		}
		nums[high] = nums[low]
	}
	nums[low] = key
	return low
}

