package goden0101

// 解法4
func isUnique(astr string) bool {
	var bin int64 = 0
	for _, b := range astr {
		dt := b - 'a'
		if bin&(1<<dt) > 0 {
			return false
		}

		bin = bin | (1 << dt)
	}

	return true
}
