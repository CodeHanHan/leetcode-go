package offer_11_s

func minArray(numbers []int) int {
	start := 0
	end := len(numbers) - 1

	for start < end {
		mid := start + (end-start)/2
		if numbers[mid] < numbers[end] {
			end = mid
		} else if numbers[mid] > numbers[end] {
			start = mid + 1
		} else {
			end = end - 1
		}
	}
	return numbers[start]
}
