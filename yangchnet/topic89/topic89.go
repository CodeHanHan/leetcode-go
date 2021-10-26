package topic89

import "math"

func grayCode(n int) []int {
	m := int(math.Pow(float64(2), float64(n)))
	ret := make([]int, 0)
	for i := 0; i < m; i++ {
		g := i ^ (i / 2)
		ret = append(ret, g)
	}
	return ret
}
