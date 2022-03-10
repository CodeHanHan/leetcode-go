package lc022

func generateParenthesis(n int) []string {
	res := []string{}

	var backtrack func(left, right int, item string)
	backtrack = func(left, right int, item string) {
		if left == right && left == n {
			res = append(res, item)
			return
		}
		if right > left || left > n || right > n {
			return
		}
		backtrack(left+1, right, item+"(")
		backtrack(left, right+1, item+")")
	}

	backtrack(0, 0, "")
	return res
}
