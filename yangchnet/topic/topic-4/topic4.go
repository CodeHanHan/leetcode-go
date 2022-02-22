package topic4

import (
	"fmt"
	"strconv"
)

// 归并法
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := make([]int, len(nums1)+len(nums2))

	var i, j int
	var p int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			nums[p] = nums1[i]
			i++
		} else {
			nums[p] = nums2[j]
			j++
		}
		p++
	}

	if i < len(nums1) {
		nums = append(nums[:p], nums1[i:]...)
	}

	if j < len(nums2) {
		nums = append(nums[:p], nums2[j:]...)
	}

	var mid float64
	if len(nums)%2 == 0 {
		mid = float64(nums[len(nums)/2]+nums[len(nums)/2-1]) / 2
	} else {
		mid = float64(nums[len(nums)/2])
	}
	f, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", mid), 64)

	return f
}

// 双指针法
func findMedianSortedArrays_1(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	var pre, cur int = -1, -1
	var p, q int = 0, 0
	for i := 0; i <= (m+n)/2; i++ {
		pre = cur
		if p < m && (q >= n || nums1[p] < nums2[q]) {
			cur = nums1[p]
			p++
		} else {
			cur = nums2[q]
			q++
		}
	}

	var mid float64
	if (m+n)&1 == 0 { // 判断奇偶的简单方法
		mid = float64(pre+cur) / 2
	} else {
		mid = float64(cur)
	}
	f, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", mid), 64)
	return f
}

// 二分法
func findMedianSortedArrays_2(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	odd := (m+n)&1 == 0
	if !odd {
		return float64(kthSmallElement(nums1, nums2, (m+n)/2+1))
	} else {
		return float64(kthSmallElement(nums1, nums2, (m+n)/2)+kthSmallElement(nums1, nums2, (m+n)/2+1)) / 2
	}
}

// getKthElement find Kth small element in combination of num1 and num2
func kthSmallElement(num1, num2 []int, k int) int {
	l1, l2 := num1, num2
	var min1, min2 int   // 存储要比较的两个最小值
	var step1, step2 int // 存储要截断的长度

	for {
		if len(l1) == 0 {
			return l2[k-1]
		}
		if len(l2) == 0 {
			return l1[k-1]
		}

		if k == 1 {
			return min(l1[0], l2[0])
		}
		half := k / 2

		// if len(l1)-1 < half-1 { // 如果l1下标将要越界
		// 	min1 = l1[len(l1)-1]
		// 	step1 = len(l1)
		// } else {
		// 	min1 = l1[half-1]
		// 	step1 = half
		// }
		// 下面两行代码从上面注释掉的代码翻译而来
		min1 = l1[min(len(l1), half)-1]
		step1 = min(len(l1), half)

		// if len(l2)-1 < half-1 { // 如果l2下标将要越界
		// 	min2 = l2[len(l2)-1]
		// 	step2 = len(l2)
		// } else {
		// 	min2 = l2[half-1]
		// 	step2 = half
		// }
		// 下面两行代码从上面注释掉的代码翻译而来
		min2 = l2[min(len(l2), half)-1]
		step2 = min(len(l2), half)

		if min1 <= min2 {
			l1 = l1[step1:] // 这里如果用append, 将会改变原数组，从而对第二次查找造成影响
			k -= step1
		} else {
			l2 = l2[step2:] // 不改变原数组，只改变指针
			k -= step2
		}
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

//TODO: 分割法
