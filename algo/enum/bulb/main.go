package main

import "fmt"

func main() {
	input := make([][]int, 6)
	for i := range input {
		input[i] = make([]int, 8)
	}
	for i := 1; i < 6; i++ {
		fmt.Scan(&input[i][1], &input[i][2], &input[i][3], &input[i][4], &input[i][5], &input[i][6])
	}

	enum := 1
	for ; enum < 64; enum++ {
		if press, ok := valid(input, enum); ok {
			for i := 1; i < 6; i++ {
				fmt.Printf("%v\n", press[i][1:len(press[i])-1])
			}
			break
		}
	}

	return
}

func valid(bulbMatrix [][]int, enum int) ([][]int, bool) {
	press := make([][]int, 6)
	for i := range press {
		press[i] = make([]int, 8)
	}

	press[1][1], press[1][2], press[1][3], press[1][4], press[1][5], press[1][6] =
		enum>>0&0x1, enum>>1&0x1, enum>>2&0x1, enum>>3&0x1, enum>>4&0x1, enum>>5&0x1

	for r := 1; r < 5; r++ {
		for c := 1; c < 7; c++ {
			// 根据press第一行以及bulbMatrix， 计算其他行的值
			press[r+1][c] = (bulbMatrix[r][c] + press[r][c] +
				press[r-1][c] + press[r][c-1] + press[r][c+1]) % 2
		}
	}

	for c := 1; c < 7; c++ {
		// 检查最后以行是否全部被关闭
		if (press[5][c-1]+press[5][c]+press[5][c+1]+press[4][c])%2 != bulbMatrix[5][c] {
			return nil, false
		}
	}

	return press, true
}
