package topic4

import (
	"fmt"
	"strconv"
)

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

// func findMedianSortedArrays_1(nums1 []int, nums2 []int) float64 {
// 	var loc int
// 	odd := (len(nums1)+len(nums2))%2 == 0
// 	if odd {
// 		loc = (len(nums1)+len(nums2))/2 - 1
// 	} else {
// 		loc = (len(nums1) + len(nums2)) / 2
// 	}

// 	var i, j, p int
// 	for i < len(nums1) && j < len(nums2) {
// 		if p >= loc {

// 		}

// 		if nums1[i] <= nums2[j] {
// 			i++
// 		}else {
// 			j++
// 		}
// 		p++
// 	}

// 	for i <len(nums1) {

// 	}

// }
