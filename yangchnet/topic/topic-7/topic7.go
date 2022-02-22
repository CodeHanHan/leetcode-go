package topic7

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	nagative := x < 0 // 是否为负数
	a := make([]int, 0)

	y := abs(x)
	for y != 0 {
		a = append(a, y%10)
		y /= 10
	}
	b := 0
	for i, v := range a {
		b += v * int(math.Pow10(len(a)-i-1))
	}

	if nagative {
		b = -b
	}

	if b > (1<<31)-1 || b < -(1<<31) {
		return 0
	}

	return b
}

func reverse_1(x int) int {
	nagative := x < 0 // 是否为负数
	y := abs(x)

	l := len(strconv.Itoa(y))

	i := 0
	b := 0

	for y != 0 {
		a := y % 10
		y = y / 10
		b += a * int(math.Pow10(l-i-1))
		i += 1
	}

	if nagative {
		b = -b
	}

	if b > (1<<31)-1 || b < -(1<<31) {
		return 0
	}

	return b

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//TODO: math.Abs(-999999999999999999) 会变成1000000000000000000
