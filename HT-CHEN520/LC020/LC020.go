package LC020

func isValid(s string) bool {
	pairs := map[byte]byte{'(': ')', '[': ']', '{': '}'}
	if len(s) == 0 {
		return true
	}

	stack := make([]byte, 0)
	for i, _ := range s {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, pairs[s[i]])
		} else if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
