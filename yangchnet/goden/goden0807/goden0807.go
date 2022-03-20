package goden0807

func permutation(S string) []string {
	if len(S) <= 0 {
		return []string{""}
	}
	res := make([]string, 0)
	used := make(map[byte]bool)

	var fn func([]byte, map[byte]bool)
	fn = func(path []byte, used map[byte]bool) {
		if len(path) == len(S) {
			tmp := make([]byte, len(path))
			copy(tmp, path)
			res = append(res, string(tmp))
			return
		}

		for i := 0; i < len(S); i++ {
			if !used[S[i]] {
				path = append(path, S[i])
				used[S[i]] = true
				fn(path, used)
				path = path[:len(path)-1]
				used[S[i]] = false
			}
		}
	}

	fn([]byte{}, used)

	return res
}
