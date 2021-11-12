package lc137

func singleNumber(nums []int) int {
	m := make(map[int]int)

	for _, k := range nums {
		m[k]++
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}
