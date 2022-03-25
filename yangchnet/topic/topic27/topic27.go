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

func removeElement2(nums []int, val int) int {
	i := 0
	for {
		if i >= len(nums) {
			break
		}

		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			continue
		}

		i++
	}

	return len(nums)
}
