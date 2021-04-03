package offer_42_s_

func maxSubArray(nums []int) int {
	var dp int = nums[0]
	var max int = dp
	for i := 1; i < len(nums); i++ {
		if dp > 0 {
			dp = dp + nums[i]
		} else {
			dp = nums[i]
		}
		if max < dp {
			max = dp
		}
	}
	return max
}
