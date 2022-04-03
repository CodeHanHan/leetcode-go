package goden1610

import "sort"

func maxAliveYear(birth []int, death []int) int {
	sort.Ints(birth)
	sort.Ints(death)
	y := birth[0]
	var curAlive, maxAlive = 0, 0
	var deathy int = 0
	for i := 0; i < len(birth); i++ {
		curAlive++
		for death[deathy] < birth[i] {
			deathy++
			curAlive--
		}
		if curAlive > maxAlive {
			y = birth[i]
			maxAlive = curAlive
		}
	}

	return y
}

func maxAliveYear2(birth []int, death []int) int {
	data, ans := make([]int, 102), 0
	for _, x := range birth {
		data[x-1900]++
	}
	for _, x := range death {
		data[x-1899]--
	}

	for i := 1; i < len(data); i++ {
		data[i] += data[i-1]
		if data[i] > data[ans] {
			ans = i
		}
	}
	return ans + 1900
}
