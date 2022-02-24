package offer05

func replaceSpace(s string) string {
	b := []byte(s)
	spaceCnt := 0
	for _, v := range s {
		if v == ' ' {
			spaceCnt++
		}
	}

	tmp := make([]byte, spaceCnt*2)
	b = append(b, tmp...)
	i, j := len(s)-1, len(b)-1
	for i >= 0 {
		if b[i] != ' ' {
			b[j] = b[i]
			i--
			j--
		} else {
			b[j] = '0'
			b[j-1] = '2'
			b[j-2] = '%'
			i--
			j -= 3
		}
	}
	
	return string(b)
}
