package topic11

// 暴力法， 超时
func maxArea(height []int) int {
	var max int = -1

	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			v := (j - i) * min(height[i], height[j])
			if max < v {
				max = v
			}
		}
	}

	return max
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// 双指针法
func maxArea_1(height []int) int {
	var i, j int = 0, len(height) - 1
	var max int = -1
	for i < j {
		if max < (j-i)*min(height[i], height[j]) {
			max = (j - i) * min(height[i], height[j])
		}
		if height[i] <= height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}
