package offer_57_II_s

func findContinuousSequence(target int) [][]int {
	res := make([][]int, 0)
	i := 1
	j := i + 1
	sum := i + j
	r := []int{i, j}
	for {
		if i*2 > target {
			break
		}
		if sum < target {
			j++
			sum += j
			r = append(r, j)
		} else if sum > target {
			i++
			j = i + 1
			sum = i + j
			r = append([]int{}, []int{i, j}...)
		} else {
			res = append(res, r)

			i++
			j = i + 1
			sum = i + j
			r = append([]int{}, []int{i, j}...)
		}
	}
	return res
}
