package offer_51_d_

func reversePairs(nums []int) int {
	l := len(nums)
	if l < 2 {
		return 0
	}
	numsCopy := make([]int, l)
	copy(numsCopy, nums)

	temp := make([]int, l)
	return reversepairs(numsCopy, 0, l-1, temp)
}

func reversepairs(nums []int, left, right int, temp []int) int {
	if left == right {
		return 0
	}

	mid := left + (right - left) / 2
	leftPairs := reversepairs(nums, left, mid, temp)
	rightPairs := reversepairs(nums, mid+1, right, temp)

	if nums[mid] <= nums[mid+1] {
		return leftPairs + rightPairs
	}

	crossPairs := mergeAndCount(nums, left, mid, right, temp)
	return leftPairs + rightPairs + crossPairs

}

func mergeAndCount(nums []int, left, mid, right int, temp []int) int {
	for i := left; i <=right; i++ {
		temp[i] = nums[i]
	}

	var i, j int = left, mid+1
	count := 0
	for k := left; k <= right; k++ {
		if i == mid + 1{		// 前半部分用尽
			nums[k] = temp[j]
			j++
		}else if j == right + 1 {	// 后半部分用尽
			nums[k] = temp[i]
			i++
		}else if temp[i] <= temp[j] {   // 如果这里是小于，则归并排序不稳定
			nums[k] = temp[i]
			i++
		}else {
			nums[k] = temp[j]
			j++
			count += mid - i + 1
		}
	}
	return count
}
