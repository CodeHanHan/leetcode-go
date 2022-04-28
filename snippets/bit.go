package snippets

// 位运算相关
// TODO: test

// 按位清零
func BitwiseClear(a *int, x int) {
	(*a) &= (0x1 << x)
}

// 按位赋1
func BitwiseTo1(a *int, x int) {
	(*a) |= (0x1 << x)
}

// 判断奇偶
func isOdd(x int) bool {
	return (x & 0x1) == 0
}

// 取某一位的值
func PickBit(a int, x int) uint8 {
	return uint8((a >> x) & 0x1)
}

// 反转位，如对uint16反转前8位
func ReversionBit(a *uint16) {
	(*a) ^= 0xFF00
}

// 检查符号是否相同
func SameSign(a, b int) bool {
	return (a ^ b) >= 0
}

// n&(n-1), 可将最右边的二进制1设为0
// 1. 检测 2 的幂
// 2. 检测一个整数二进制中 1 的个数
// 3. 一个数变成另一个数其中改变了多少个bit位

// 使用 n&(n-1) 计算二进制1的个数, 可参考下图
// ![20220428104951](https://raw.githubusercontent.com/lich-Img/blogImg/master/img/20220428104951.png)
func CountBit1(a int) uint8 {
	var c uint8
	for a > 0 {
		c++
		a &= a - 1
	}

	return c
}

// 使用n&(-n)获取二进制最后一位1
// 应用：判断一个数是否是2的幂
// func solution(x int) bool {
// 		return n>0&(n&(-n))==n
// }
func Last1(a int) int {
	return a & (-a)
}

// 两个技巧
// 1. “异或”是一个无进位加法，说白了就是把进位砍掉。比如01^01=00。
// 2. “与”可以用来获取进位，比如01&01=01，然后再把结果左移一位，就可以获取进位结果。

// 位运算求相反数
func reversal(a int) int {
	return ^a + 1 // ^做一元操作符时为按位取反
	// a = 0111 = 7, ^a = 1000, -a+1 = 1001 = -7
}

// 位运算求绝对值
func absByBit(a int) int {
	i := a >> 31
	if i == 0 { // 正数
		return a
	}
	return ^a + 1
}

// 高低位交换
func swapLowAndHigh(a int16) int16 {
	a = (a >> 8) | (a << 8)
	return a
}
