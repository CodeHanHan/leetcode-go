package topic275

import "sort"

func hIndex(citations []int) int {
	n := len(citations)
	return n - sort.Search(n, func(x int) bool { return citations[x] >= n-x })
}
