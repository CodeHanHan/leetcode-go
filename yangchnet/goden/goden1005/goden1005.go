package goden1005

func findString(words []string, s string) int {
	l, r := 0, len(words)-1

	var mid, lmid int
	for l <= r {
		mid = l + (r-l)/2

		lmid = mid
		if words[lmid] == "" {
			for words[lmid] == "" {
				lmid--
			}
		}

		if words[lmid] == s {
			return lmid
		} else if words[lmid] < s {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}
