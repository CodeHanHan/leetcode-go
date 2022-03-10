package backtrack

/**
* 正则表达式
* 假设正则表达式中只包含“*”和“?”这两种通配符，并且对这两个通配符的语义稍微做些改变
* 其中，“*”匹配任意多个（大于等于 0 个）任意字符，“?”匹配零个或者一个任意字符
* 基于以上背景假设，如何用回溯算法，判断一个给定的文本，能否跟给定的正则表达式匹配
**/

type Regexp struct {
	Pattern string // 正则表达式
}

func (re *Regexp) Match(s string) bool {
	matched := false

	var fn func(si, pi int, s string) // si: 字符下标; pi: 正则下标;
	fn = func(si, pi int, s string) {
		if matched {
			return
		}

		if pi == len(re.Pattern) {
			if si == len(s) {
				matched = true
			}
			return
		}

		if re.Pattern[pi] == '*' { // * 匹配任意个字符
			for k := 0; k <= len(s)-si; k++ {
				fn(si+k, pi+1, s)
			}
		} else if re.Pattern[pi] == '?' { // ? 匹配0个或1个字符
			fn(si, pi+1, s)
			fn(si+1, pi+1, s)
		} else if si < len(s) && re.Pattern[pi] == s[si] { // 纯字符匹配
			fn(si+1, pi+1, s)
		}
	}

	fn(0, 0, s)

	return matched
}
