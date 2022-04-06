package goden1620

func getValidT9Words(num string, words []string) []string {
	m := map[byte]string{
		'1': "!@#",
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	candidate := make(map[string]int)
	var dfs func(i int, can []byte)
	dfs = func(i int, can []byte) {
		if i == len(num) {
			candidate[string(can)] = 1
			return
		}

		for j := 0; j < len(m[num[i]]); j++ {
			dfs(i+1, append(can, m[num[i]][j]))
		}
	}

	dfs(0, []byte{})

	ret := make([]string, 0)
	for _, word := range words {
		if candidate[word] == 1 {
			ret = append(ret, word)
		}
	}

	return ret
}

func getValidT9Words2(num string, words []string) []string {
	m := map[byte]byte{'a': '2', 'b': '2', 'c': '2', 'd': '3', 'e': '3', 'f': '3',
		'g': '4', 'h': '4', 'i': '4', 'j': '5', 'k': '5', 'l': '5',
		'm': '6', 'n': '6', 'o': '6', 'p': '7', 'q': '7', 'r': '7', 's': '7',
		't': '8', 'u': '8', 'v': '8', 'w': '9', 'x': '9', 'y': '9', 'z': '9'}

	ret := make([]string, 0)
	for _, word := range words {
		if len(word) != len(num) {
			continue
		}

		i := 0
		for ; i < len(word); i++ {
			if m[word[i]] != num[i] {
				break
			}
		}

		if i >= len(word) {
			ret = append(ret, word)
		}
	}

	return ret
}
