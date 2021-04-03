package offer_40_s_

func getLeastNumbers(arr []int, k int) []int {
	if k <= 0 {
		return []int{}
	}
	minNums := make([]int, k)
	minNums = arr[:k]
	max, maxI := getMax(minNums)

	for i := k; i < len(arr); i++ {
		if arr[i] < max {
			minNums[maxI] = arr[i]
			max, maxI = getMax(minNums)
		}
	}
	return minNums
}

func getMax(arr []int) (max, maxI int) {
	max = arr[0]
	maxI = 0
	for i := 1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
			maxI = i
		}
	}
	return max, maxI
}
