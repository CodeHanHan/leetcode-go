package topic287

// 如果只有一个数字重复了一次
// func findDuplicate(nums []int) int {
// 	n := len(nums) - 1
// 	m := nums[0]
// 	for i := 1; i < len(nums); i++ {
// 		m ^= nums[i]
// 	}

// 	for i := 1; i <= n; i++ {
// 		m ^= i
// 	}

// 	return m
// }

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		fast = nums[nums[fast]]
		slow = nums[slow]
		if slow == fast {
			fast = 0
			for nums[slow] != nums[fast] {
				fast = nums[fast]
				slow = nums[slow]
			}

			return nums[slow]
		}
	}
}
