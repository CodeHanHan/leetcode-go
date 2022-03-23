package goden0809

func generateParenthesis(n int) []string {
	ret := make([]string, 0)

	var fn func(l, r int, gen string)
	fn = func(l, r int, gen string) {
		if l == 0 && r == 0 {
			ret = append(ret, gen)
			return
		}

		if l > 0 {
			gen += "("
			fn(l-1, r, gen)
			gen = gen[:len(gen)-1]
		}

		if r > l {
			gen += ")"
			fn(l, r-1, gen)
			gen = gen[:len(gen)-1]
		}
	}

	fn(n, n, "")

	return ret
}
