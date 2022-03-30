package lc350

import "sort"

func intersect1(nums1 []int, nums2 []int) []int {
	var res []int
	sort.Ints(nums1)
	sort.Ints(nums2)

	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}
	return res
}

func intersect2(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect2(nums2, nums1)
	}

	m := map[int]int{}
	for _, num := range nums1 {
		m[num]++
	}
	res := []int{}
	for _, num := range nums2 {
		if m[num] > 0 {
			res = append(res, num)
			m[num]--
		}
	}
	return res
}
