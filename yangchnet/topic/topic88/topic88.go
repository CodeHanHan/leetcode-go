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

func merge1(nums1 []int, m int, nums2 []int, n int) {
	p, q := m-1, n-1
	v := m + n - 1
	for q >= 0 {
		if p < 0 || nums1[p] <= nums2[q] {
			nums1[v] = nums2[q]
			q--
		} else {
			nums1[v] = nums1[p]
			p--
		}
		v--
	}
}
