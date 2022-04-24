package lc287

func findDuplicate1(nums []int) int {
	n := len(nums)
	left, right := 1, n-1
	for left < right {
		cnt := 0
		mid := (left + right) / 2
		for i := 0; i < n; i++ {
			if nums[i] <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func findDuplicate2(nums []int) int {
	slow, fast := 0, 0
	for{
		slow=nums[slow]
		fast=nums[nums[fast]]

		if slow==fast {
			fast=0
			for{
				if slow==fast {
					return slow
				}
				slow=nums[slow]
				fast=nums[fast]
			}
		}
	}
}
