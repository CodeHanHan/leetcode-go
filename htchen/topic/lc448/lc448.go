package lc448

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	var res []int
	for _, v := range nums {
		v = (v - 1) % n
		nums[v] += n
	}
	for i, v := range nums {
		if v <= n {
			res = append(res, i+1)
		}
	}
	return res
}
