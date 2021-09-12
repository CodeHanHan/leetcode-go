package offer_51_d_


func MergeSort(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}
	numsCopy := make([]int, l)
	copy(numsCopy, nums)

	temp := make([]int, l)

	sort(numsCopy, 0, l-1, temp)
	return numsCopy
}

func sort(nums []int, left, right int, temp []int)  {
	if left == right {return }

	mid := left + (right - left) / 2  //防止left和right过大
	sort(nums, left, mid, temp)
	sort(nums, mid+1, right, temp)

	if nums[mid] <= nums[mid+1] {return}

	merge(nums, left, mid, right, temp)

}

func merge(nums []int, left, mid, right int, temp []int) {
	for i := left; i <= right; i++ {
		temp[i] = nums[i]
	}

	var i, j int = left, mid+1

	for k := left; k <= right; k++ {
		if i == mid + 1 {
			nums[k] = temp[j]
			j++
		} else if j == right + 1 {
			nums[k] = temp[i]
			i++
		} else if temp[i] <= temp[j] {
			nums[k] = temp[i]
			i++
		}else {
			nums[k] = temp[j]
			j++
		}
	}
}
