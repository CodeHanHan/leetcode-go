package goden1717

func multiSearch(big string, smalls []string) [][]int {
	length := make(map[int]bool)
	words := make(map[string]int, 0)
	for i, small := range smalls {
		if small != "" {
			length[len(small)] = true
			words[small] = i
		}
	}

	ret := make([][]int, len(smalls))
	for i := range ret {
		ret[i] = make([]int, 0)
	}

	for i := 0; i < len(big); i++ {
		for k, _ := range length {
			if i+k > len(big) {
				continue
			}

			if idx, ok := words[big[i:i+k]]; ok {
				ret[idx] = append(ret[idx], i)
			}
		}
	}

	return ret
}
