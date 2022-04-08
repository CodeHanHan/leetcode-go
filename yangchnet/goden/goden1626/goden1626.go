package goden1626

func calculate(s string) int {
	numStack := make([]int, 0)
	expStack := make([]byte, 0)

	p := map[byte]int{
		'*': 2,
		'/': 2,
		'+': 1,
		'-': 1,
	}

	for i := 0; i < len(s); {
		if s[i] == ' ' {
			i++
			continue
		}

		if s[i] == '*' || s[i] == '/' || s[i] == '+' || s[i] == '-' {
			for len(expStack) > 0 && p[s[i]] <= p[expStack[len(expStack)-1]] { // 当前符号优先级小于栈顶运算符优先级
				exp := expStack[len(expStack)-1]
				expStack = expStack[:len(expStack)-1]

				n1 := numStack[len(numStack)-2]
				n2 := numStack[len(numStack)-1]

				numStack = numStack[:len(numStack)-2]
				numStack = append(numStack, cal(n1, n2, exp))
			}
			expStack = append(expStack, s[i])
			i++
			continue
		}

		num := 0
		j := i
		for j < len(s) && s[j] != ' ' && s[j] != '*' && s[j] != '/' && s[j] != '+' && s[j] != '-' {
			num = num*10 + int(s[j]-'0')
			j++
		}

		numStack = append(numStack, num)
		i = j
	}

	for len(expStack) > 0 {
		exp := expStack[len(expStack)-1]
		expStack = expStack[:len(expStack)-1]

		n1 := numStack[len(numStack)-2]
		n2 := numStack[len(numStack)-1]

		numStack = numStack[:len(numStack)-2]
		numStack = append(numStack, cal(n1, n2, exp))
	}

	return numStack[len(numStack)-1]
}

func cal(n1, n2 int, exp byte) int {
	switch exp {
	case '*':
		return n1 * n2
	case '/':
		return n1 / n2
	case '+':
		return n1 + n2
	case '-':
		return n1 - n2
	}

	return 0
}
