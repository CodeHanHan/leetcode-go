package topic137

import (
	"sort"
)

func singleNumber(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ { // 第31位为符号位
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}
		if total%3 > 0 {
			ans |= 1 << i
		}
	}

	return int(ans)
}

func singleNumber_1(nums []int) int {
	sort.Ints(nums)

	for i := 0; i < len(nums); i += 3 {
		if i+1 >= len(nums) {
			return nums[i]
		}
		if d, diff := different(nums[i], nums[i+1], nums[i+2]); diff {
			return d
		}
	}

	return 0
}

func different(a, b, c int) (int, bool) { // bool 为false代表三个数相同
	if a == b {
		if b == c {
			return 0, false
		} else {
			return c, true
		}
	} else if b == c {
		return a, true
	} else if a == c {
		return b, true
	}
	return 0, false
}
