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

// 栈
func longestValidParentheses1(s string) int {
	maxAns := 0
	stack := []int{}
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxAns = max(maxAns, i-stack[len(stack)-1])
			}
		}
	}
	return maxAns
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 动态规划
func longestValidParentheses2(s string) int {
	maxAns := 0
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
			maxAns = max(maxAns, dp[i])
		}
	}
	return maxAns
}
