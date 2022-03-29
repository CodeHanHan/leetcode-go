package lc229

//hash统计
func majorityElement1(nums []int) []int {
	cnt := map[int]int{}
	for _, v := range nums {
		cnt[v]++
	}
	res := make([]int, 0)
	for v, c := range cnt {
		if c > len(nums)/3 {
			res = append(res, v)
		}
	}
	return res
}

//摩尔投票法
func majorityElement2(nums []int) []int {
	n := len(nums)
	var res []int
	element1, element2 := 0, 0
	vote1, vote2 := 0, 0

	for _, v := range nums {
		if vote1 > 0 && v == element1 {
			vote1++
		} else if vote2 > 0 && v == element2 {
			vote2++
		} else if vote1 == 0 {
			element1 = v
			vote1++
		} else if vote2 == 0 {
			element2 = v
			vote2++
		} else {
			vote1--
			vote2--
		}
	}

	cnt1, cnt2 := 0, 0
	for _, v := range nums {
		if vote1 > 0 && v == element1 {
			cnt1++
		}
		if vote2 > 0 && v == element2 {
			cnt2++
		}
	}

	if vote1 > 0 && cnt1 > n/3 {
		res = append(res, element1)
	}
	if vote2 > 0 && cnt2 > n/3 {
		res = append(res, element2)
	}

	return res
}
