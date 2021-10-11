package topic26

import "fmt"

func removeDuplicates(nums []int) int {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	fmt.Printf("%v", nums)
	return len(nums)
}
