package topic118

func generate(numRows int) [][]int {
	ans := make([][]int, 0)
	ans = append(ans, []int{1})
	if numRows == 1 {
		return ans
	}
	ans = append(ans, []int{1, 1})
	if numRows == 2 {
		return ans
	}

	for i := 3; i <= numRows; i++ {
		tmp := make([]int, i)
		tmp[0], tmp[i-1] = 1, 1
		for j := 1; j < i-1; j++ {
			tmp[j] = ans[i-2][j-1] + ans[i-2][j]
		}
		ans = append(ans, tmp)

	}

	return ans
}
