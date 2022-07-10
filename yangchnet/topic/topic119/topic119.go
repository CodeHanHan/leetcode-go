package topic119

func getRow(rowIndex int) []int {
	ans := make([]int, rowIndex+1)
	for i, _ := range ans {
		ans[i] = 1
	}

	if rowIndex <= 1 {
		return ans[:rowIndex+1]
	}

	for i := 3; i <= rowIndex+1; i++ {
		for j := i - 2; j > 0; j-- {
			ans[j] = ans[j-1] + ans[j]
		}
	}

	return ans
}
