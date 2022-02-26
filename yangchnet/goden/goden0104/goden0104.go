package goden0104

func canPermutePalindrome(s string) bool {
	count := make(map[rune]int)
	for _, c := range s {
		count[c] += 1
	}

	odd := 0
	for _, c := range count {
		if c%2 != 0 {
			odd++
		}
		if odd > 1 {
			return false
		}
	}

	return true
}
