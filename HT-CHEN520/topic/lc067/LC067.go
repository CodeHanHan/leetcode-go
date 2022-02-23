package lc067

func addBinary(a string, b string) string {
	la := len(a)
	lb := len(b)
	if la < lb {
		return add(a, b, la, lb)
	} else {
		return add(b, a, lb, la)
	}
}

func add(a, b string, la, lb int) string {

	for i := 0; i < lb-la; i++ {
		a = "0" + a
	}

	flag := 0
	ret := []byte(a)
	for i := lb - 1; i >= 0; i-- {
		ret[i] = ret[i] + b[i] - byte('0') + byte(flag)
		if ret[i] >= byte('2') {
			ret[i] -= 2
			flag = 1
			continue
		}
		flag = 0
	}

	if flag == 1 {
		ret = append([]byte{'1'}, ret...)
	}

	return string(ret)
}
