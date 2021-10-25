package topic88

func merge(nums1 []int, m int, nums2 []int, n int) {
	tmp := make([]int, 0)

	var i, j int = 0, 0
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			tmp = append(tmp, nums1[i])
			i++
		} else {
			tmp = append(tmp, nums2[j])
			j++
		}
	}

	for ; i < m; i++ {
		tmp = append(tmp, nums1[i])
	}

	for ; j < n; j++ {
		tmp = append(tmp, nums2[j])
	}

	copy(nums1, tmp)
}
