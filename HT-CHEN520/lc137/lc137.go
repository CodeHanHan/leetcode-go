package lc137

func singleNumber(nums []int) int {
	m := make(map[int]int)

	for _, k := range nums {
		m[k]++
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}

func singleNumber_1(nums []int) int {
	bitNum, res := int32(0), int32(0)
	for i := 0; i < 32; i++ {
		bitNum = int32(0)

		for _, val := range nums {
			bitNum += (int32(val) >> i) & 1
		}
		res |= (bitNum % 3) << i
	}
	return int(res)
}
