package goden1722

func findLadders(beginWord string, endWord string, wordList []string) []string {
	// endWord是否在wordList中，若在，找寻其位置
	target := findWord(endWord, wordList)
	if target < 0 {
		return []string{}
	}

	// beginWord是否在wordList中，若在，找寻其位置
	start := findWord(beginWord, wordList)
	if start < 0 {
		wordList = append([]string{beginWord}, wordList...)
		target += 1
		start = 0
	}

	n := len(wordList)

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	// 如果wordList[i] 与 wordList[j]只有一个字符不同，则matrix[i][j] = 1, 意为二者可达
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && diff(wordList[i], wordList[j]) {
				matrix[i][j] = 1
			}
		}
	}

	paths := make([][]int, 0)
	visited := make([]bool, n)
	visited[target] = true

	findExistsPath([]int{target}, &paths, visited, matrix, start, target)

	// 没找到有效路径
	if len(paths) <= 0 {
		return []string{}
	}

	ret := make([]string, 0)
	for i := len(paths[0]) - 1; i >= 0; i-- {
		ret = append(ret, wordList[paths[0][i]])
	}

	return ret
}

// 在words中定位word
func findWord(word string, words []string) int {
	for i := 0; i < len(words); i++ {
		if word == words[i] {
			return i
		}
	}

	return -1
}

// 找到有效的路径, 这个函数参考了 “面试题 04.01.节点间通路”
func findExistsPath(path []int, paths *[][]int, visited []bool, matrix [][]int, start, target int) {
	if start == target {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*paths = append(*paths, tmp)
		return
	}

	for i, m := range matrix[target] {
		if m == 1 {
			if visited[i] {
				continue
			}
			visited[i] = true
			findExistsPath(append(path, i), paths, visited, matrix, start, i)
		}
	}

	return
}

// 两个单词是否只有一处不同
func diff(word1, word2 string) bool {
	c := 0
	for i := 0; i < len(word1); i++ {
		if word1[i] != word2[i] {
			c++
		}
		if c > 1 {
			return false
		}
	}
	return c == 1
}
