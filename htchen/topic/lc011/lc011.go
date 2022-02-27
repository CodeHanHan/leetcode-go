package lc011

func maxArea(height []int) int {
	maxWater := 0
	i, j := 0, len(height)-1
	for i < j {
		tempRes := min(height[i], height[j]) * (j - i)
		if tempRes > maxWater {
			maxWater = tempRes
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxWater
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
