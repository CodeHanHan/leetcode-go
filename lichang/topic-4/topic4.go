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

//TODO: 二分法

//TODO: 分割法
