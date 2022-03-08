package lc017

var phoneNum map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	conbinations := []string{}

	var backtrack func(digits string, n int, conbination string)
	backtrack = func(digits string, n int, conbination string) {
		if n == len(digits) {
			conbinations = append(conbinations, conbination)
			return
		}
		digit := string(digits[n])
		letters := phoneNum[digit]
		lettersNums := len(letters)
		for i := 0; i < lettersNums; i++ {
			backtrack(digits, n+1, conbination+string(letters[i]))
		}
	}

	backtrack(digits, 0, "")
	return conbinations
}
