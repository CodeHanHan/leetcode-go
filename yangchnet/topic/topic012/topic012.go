package topic012

func intToRoman(num int) string {
	var roman map[int]string = map[int]string{
		1000: "M",
		500:  "D",
		100:  "C",
		50:   "L",
		10:   "X",
		5:    "V",
		1:    "I",
	}

	var romanNums []int = []int{1000, 500, 100, 50, 10, 5, 1}

	var ret string = ""
	var pre int = -1
	for i := 0; i < len(romanNums); i++ {
		t := num / romanNums[i]
		if t == 4 { // 是4或9
			if pre == 500 || pre == 50 || pre == 5 { // 是9
				ret = ret[:len(ret)-1]
				ret += roman[romanNums[i]]
				ret += roman[romanNums[i-1]*2]
			} else { // 是4
				ret += roman[romanNums[i]]
				ret += roman[romanNums[i-1]]

			}
		} else {
			for j := 0; j < t; j++ {
				ret += roman[romanNums[i]]
			}
		}

		num = num - t*romanNums[i]
		pre = t * romanNums[i]
	}

	return ret
}

func intToRoman2(num int) string {
	var valueSymbols = []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	roman := []byte{}
	for _, vs := range valueSymbols {
		for num >= vs.value {
			num -= vs.value
			roman = append(roman, vs.symbol...)
		}
		if num == 0 {
			break
		}
	}
	return string(roman)
}
