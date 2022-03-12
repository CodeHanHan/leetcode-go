package goden0503

func reverseBits(num int) int {
	cur := 0
	insert := 0
	res := 1
	for i := 0; i < 32; i++ {
		b := num >> i & 0x1
		if b == 1 {
			cur += 1
			insert += 1
		} else {
			insert = cur + 1
			cur = 0
		}

		if res < insert {
			res = insert
		}
	}

	return res
}
