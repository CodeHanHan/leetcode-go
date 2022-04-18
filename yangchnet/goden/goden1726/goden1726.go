package goden1726

import (
	"fmt"
)

// HACK 超时

// func computeSimilarities(docs [][]int) []string {
// 	for i := range docs {
// 		sort.Ints(docs[i])
// 	}

// 	ret := make([]string, 0)
// 	for i := 0; i < len(docs); i++ {
// 		for j := i + 1; j < len(docs); j++ {
// 			diff, same := diffAndSame(docs[i], docs[j])
// 			if same > 0 {
// 				ret = append(ret, fmt.Sprintf("%d,%d: %.4f", i, j, float64(same)/float64(diff+same)+1e-9))
// 			}
// 		}
// 	}

// 	return ret
// }

// func diffAndSame(arr1, arr2 []int) (diff, same int) {
// 	p, q := 0, 0

// 	diffMap := make(map[int]bool)
// 	for p < len(arr1) && q < len(arr2) {
// 		if arr1[p] > arr2[q] {
// 			// q右移
// 			diffMap[arr1[p]] = true
// 			diffMap[arr2[q]] = true
// 			q++
// 		} else if arr1[p] < arr2[q] {
// 			// p右移
// 			diffMap[arr1[p]] = true
// 			diffMap[arr2[q]] = true
// 			p++
// 		} else if arr1[p] == arr2[q] {
// 			// 均右移
// 			delete(diffMap, arr1[p])
// 			p++
// 			q++
// 			same++
// 		}
// 	}

// 	for p < len(arr1) {
// 		diffMap[arr1[p]] = true
// 		p++
// 	}

// 	for q < len(arr2) {
// 		diffMap[arr2[q]] = true
// 		q++
// 	}

// 	return len(diffMap), same
// }

func computeSimilarities(docs [][]int) []string {
	n := len(docs)
	similar := make([][]int, n) // 记录集合1和2中交集个数
	for i := range similar {
		similar[i] = make([]int, n)
	}

	// 统计单词在哪个文档中出现过
	m := make(map[int][]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < len(docs[i]); j++ {
			if len(m[docs[i][j]]) <= 0 {
				m[docs[i][j]] = make([]int, 0)
			}
			m[docs[i][j]] = append(m[docs[i][j]], i)
		}
	}

	// 若有某个单词在两个及以上文档中出现过，则标记之
	for _, v := range m {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				similar[v[i]][v[j]]++
				// similar[i][j]++ // BUG
			}
		}
	}

	// 稀疏相似性 = 集合1集合2交集个数 / (集合1集合2总个数 - 集合1集合2交集个数)
	ret := make([]string, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if similar[i][j] <= 0 {
				continue
			}

			rate := float64(similar[i][j]) / (float64(len(docs[i]) + len(docs[j]) - similar[i][j]))
			ret = append(ret,
				fmt.Sprintf("%d,%d: %.4f", i, j, rate+1e-9))
		}
	}

	return ret
}
