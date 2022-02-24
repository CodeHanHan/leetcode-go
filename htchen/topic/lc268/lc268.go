package lc268

func missingNumber(nums []int) int {
	n := len(nums)
	totalSum := n * (n + 1) / 2
	arrSum := 0
	for _, num := range nums {
		arrSum += num
	}
	return totalSum - arrSum
}

func missingNumber_1(nums []int) int {
	res := 0
	for i, v := range nums {
		res ^= i ^ v
	}
	return res ^ len(nums)
}

func missingNumber_2(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}
	for i := 0; ; i++ {
		if !m[i] {
			return i
		}
	}
}
