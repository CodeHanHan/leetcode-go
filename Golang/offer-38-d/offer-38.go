package offer_38

// 回溯法，使用bool数组来标记已使用的字符
//func permutation(s string) []string {
//	res := []string{}
//	str := []byte(s)
//	var path []byte
//	used := make([]bool, len(str))
//	dict := make(map[string]bool)
//	var f func(str []byte, used []bool, path []byte)
//
//	f = func(str []byte, used []bool, path []byte) {
//
//		if  len(path) == len(s) {
//			if dict[string(path)] {return}
//			res = append(res, string(path))
//			dict[string(path)] = true
//			return
//		}
//
//		for i, v := range str {
//			if !used[i] {
//				path = append(path, v)
//				used[i] = true
//				f(str, used, path)
//
//				// TODO: 回复状态
//				used[i] = false
//				path = path[:len(path)-1]
//			}
//		}
//	}
//	f(str, used, path)
//	return res
//}

// 牛逼解法，没看懂
func permutation(s string) []string {
	res := []string{}
	bytes := []byte(s)
	var dfs func(x int)
	dfs = func(x int) {
		if x == len(bytes)-1 {
			res = append(res, string(bytes))
		}
		dict := map[byte]bool{}
		for i := x; i < len(bytes); i++ {
			if !dict[bytes[i]] {
				bytes[x], bytes[i] = bytes[i], bytes[x]
				dict[bytes[x]] = true
				dfs(x + 1)
				bytes[x], bytes[i] = bytes[i], bytes[x]
			}
		}
	}
	dfs(0)
	return res
}
