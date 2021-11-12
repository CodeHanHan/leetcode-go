package topic013

func romanToInt(s string) int {
	roman := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}

	ret := 0
	i := 0
	for i < len(s) {
		if i+2 <= len(s) {
			if num, ok := roman[s[i:i+2]]; ok {
				ret += num
				i += 2
				continue
			}
		}
		if num, ok := roman[s[i:i+1]]; ok {
			ret += num
			i += 1
			continue
		}
	}

	return ret
}

func romanToInt_1(s string) int {
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	ret := 0
	i := len(s) - 1
	for i >= 0 {
		if i+1 < len(s) && roman[s[i]] < roman[s[i+1]] {
			ret -= roman[s[i]]
		} else {
			ret += roman[s[i]]
		}
		i--
	}

	return ret
}
