package topic470

/**古典概型
1. 第一次rand7限定[1,6]，判断奇偶性，概率是1/2
2. 第二次rand7限定[1,5]，概率是1/5
3. 二者结合可以得出10种概率相同的结果
*/

// func rand10() int {
// 	var n, m int

// 	for {
// 		n = rand7()
// 		if n <= 6 {
// 			break
// 		}
// 	}

// 	for {
// 		m = rand7()
// 		if m <= 5 {
// 			break
// 		}
// 	}

// 	if n&0x1 == 1 {
// 		return m
// 	} else {
// 		return m + 5
// 	}
// }

// func rand7() int
