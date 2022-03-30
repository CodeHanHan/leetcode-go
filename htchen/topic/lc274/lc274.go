package lc274

import (
	"sort"
)

func hIndex1(citations []int) int {
	n, h := len(citations), 0
	sort.Ints(citations)
	for i := n - 1; i >= 0 && citations[i] > h; i-- {
		h++
	}
	return h
}

func hIndex2(citations []int) int {
	n := len(citations)
	cnt:= make([]int, n+1)
	for _, citation := range citations {
		if citation >= n {
			cnt[n]++
		} else {
			cnt[citation]++
		}
	}

	for i, tot := n, 0; i >= 0; i-- {
		tot += cnt[i]
		if tot >= i {
			return i
		}
	}
	return 0
}
