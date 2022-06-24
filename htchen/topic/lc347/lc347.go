package lc347

func topKFrequent(nums []int, k int) []int {
	var ans []int
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	buckets := make([][]int, len(nums)+1)
	for k, v := range m {
		buckets[v] = append(buckets[v], k)
	}

	for i := len(buckets) - 1; k > 0; i-- {
		for _, v := range buckets[i] {
			ans = append(ans, v)
			k--
		}
	}
	return ans
}
