package topic88

func merge(nums1 []int, m int, nums2 []int, n int) {
	l := m
	k := 0
	for i := 0; i < n; i++ {
		for j := k; j < m; j++ {
			if nums2[i] > nums1[j] && nums2[i] <= nums1[j] {
				tmp := make([]int, l-j)
				copy(tmp, nums2[j:l])
				nums1 = append(nums1[:j], nums2[i])
				nums1 = append(nums1[:j+1], tmp...)
			}
		}
	}
}
