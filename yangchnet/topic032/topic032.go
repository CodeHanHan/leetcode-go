package topic032

type group struct {
	c byte
	i int
}

func longestValidParentheses(s string) int {
	match := make([]bool, len(s))
	stack := make([]*group, 0)

	for i, c := range s {
		if len(stack) > 0 && c == ')' && stack[len(stack)-1].c == '(' { // 栈不空，且栈顶元素和待入栈元素为一对括号
			match[i] = true
			match[stack[len(stack)-1].i] = true
			stack = stack[:len(stack)-1]
			continue
		} else { // 栈为空，或栈顶元素无法与待入栈元素组合
			stack = append(stack, &group{
				c: byte(c),
				i: i,
			})
		}
	}

	maxLen := 0
	for i := 0; i < len(match); {
		if match[i] {
			j := i
			for j < len(match) && match[j] {
				j++
			}
			if maxLen < j-i {
				maxLen = j - i
			}
			i = j
			continue
		}
		i++
	}

	return maxLen
}
