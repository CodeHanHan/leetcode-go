package topic018

import "sort"

func fourSum(nums []int, target int) [][]int {
	rets := make([][]int, 0)
	sort.Ints(nums)
	dfs(nums, target, []int{}, func(ret []int) {
		rets = append(rets, ret)
	})
	return rets
}

func dfs(nums []int, target int, ret []int, visit func([]int)) {
	if len(ret) == 4 && sum(ret) == target {
		r := make([]int, len(ret))
		copy(r, ret)
		visit(r)
		return
	}

	for i := 0; i < len(nums); i++ {
		if len(nums) < 4-len(ret) { // 剩余可选数字小于所需数字数目
			return
		}

		if i >= 1 && nums[i] == nums[i-1] { // 当前数字和上一个数字相同
			continue
		}

		if i+1 <= len(nums)-1 && sum(ret)+nums[i]+(4-len(ret)-1)*nums[i+1] > target { // 已选数字+当前数字+未选数字中最小的一个*待选数目>target，最后所得结果一定大于target
			return
		}

		if sum(ret)+nums[i]+(4-len(ret)-1)*nums[len(nums)-1] < target { // 已选数字+当前数字+未选数字中最大的一个*待选数目<target，最后所得结果一定小于target
			continue
		}

		ret = append(ret, nums[i])
		dfs(nums[i+1:], target, ret, visit)
		ret = ret[:len(ret)-1]
	}
}

func sum(nums []int) int {
	s := 0
	for i := 0; i < len(nums); i++ {
		s += nums[i]
	}
	return s
}
