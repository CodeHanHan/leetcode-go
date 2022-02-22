package topic27

import "fmt"

func removeElement(nums []int, val int) int {
	var p, q int = 0, 0
	var l = 0
	for q < len(nums) {
		if nums[q] != val {
			nums[p] = nums[q]
			l++
			p++
		}
		q++
	}
	nums = nums[:l]
	fmt.Printf("%v", nums)
	return l
}
