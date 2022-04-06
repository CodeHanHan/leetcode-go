package goden1618

func patternMatching(pattern string, value string) bool {
	// 1. pattern为空
	if pattern == "" {
		return len(value) == 0
	}

	// 2. pattern不为空，记录a，b的个数
	var aCount, bCount int = 0, 0
	for _, v := range pattern {
		if v == 'a' {
			aCount++
		} else {
			bCount++
		}
	}

	// 2.1 value为"", 判断patter是否为一种字符组成
	if value == "" {
		return aCount == 0 || bCount == 0
	}

	// 2.2 value不为空
	// 2.2.1 处理patter只有一种字符串的情况
	if aCount == 0 {
		return canDivide(value, bCount)
	}
	if bCount == 0 {
		return canDivide(value, aCount)
	}

	// 2.2.2 处理pattern中'a', 'b'可为""的情况
	if canDivide(value, aCount) || canDivide(value, bCount) {
		return true
	}

	// 2.2.3 枚举'a', 'b'所代表的字符串长度
	//使得aCount*aLen+bCount*bLen=valLen
	for aLen := 1; aLen*aCount <= len(value)-bCount; aLen++ {
		// value长度减去a所代表的字符串，个数不够为bCount的倍数
		if (len(value)-aLen*aCount)%bCount != 0 {
			continue
		}

		bLen := (len(value) - aLen*aCount) / bCount
		if check(pattern, value, aLen, bLen) {
			return true
		}
	}

	return false
}

// canDivide 检查value是否可被分割成相同的k份
func canDivide(value string, k int) bool {
	if len(value)%k != 0 {
		return false
	}

	d := len(value) / k
	for i := d; i < len(value); i += d {
		if value[i:i+d] != value[0:d] {
			return false
		}
	}

	return true
}

// check 检查模式是否匹配
func check(pattern, value string, aLen, bLen int) bool {
	var a, b string // 分别记录a和b匹配到的字符串

	for i, j := 0, 0; i < len(pattern); i++ {
		if pattern[i] == 'a' {
			if a == "" {
				a = value[j : j+aLen]
			} else if a != value[j:j+aLen] {
				return false
			}
			j += aLen
		} else {
			if b == "" {
				b = value[j : j+bLen]
			} else if b != value[j:j+bLen] {
				return false
			}
			j += bLen
		}
	}

	return a != b
}
