package lc027

func removeElement(nums []int, val int) int {
	len := len(nums)
	slow := 0
	for fast := 0; fast < len; fast++ {
		if nums[fast] == val {
			continue
		}
		nums[slow] = nums[fast]
		slow++
	}
	return slow
}
