package lc283

func moveZeroes(nums []int) {
	fast, slow, n := 0, 0, len(nums)
	for fast < n {
		if nums[fast] != 0 {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
		fast++
	}
}
