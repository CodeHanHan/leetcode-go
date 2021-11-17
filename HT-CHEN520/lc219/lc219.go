package lc219

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, len(nums))
	for i, val := range nums {
		if _, ok := m[val]; ok && i-m[val] <= k {
			return true
		}
		m[val] = i
	}
	return false
}
