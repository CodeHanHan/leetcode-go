package topic20

func isValid(s string) bool {
	var s1 []rune = make([]rune, 0)
	var m map[rune]rune = map[rune]rune{'(': ')', '[': ']', '{': '}'}

	for _, v := range s {
		switch v {
		case '(', '[', '{':
			s1 = append(s1, v)
		case ')', ']', '}':
			if len(s1) <= 0 || m[s1[len(s1)-1]] != v {
				return false
			}
			s1 = s1[:len(s1)-1]
		}
	}

	return len(s1) <= 0
}
