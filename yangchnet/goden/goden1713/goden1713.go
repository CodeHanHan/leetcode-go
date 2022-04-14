package goden1713

// FIXME: 需从尾部寻找单词

// func respace(dictionary []string, sentence string) int {
// 	dt := NewTrie()
// 	for _, word := range dictionary {
// 		dt.Insert(word)
// 	}

// 	dp := make([]int, len(sentence)+1)
// 	dp[0] = 0 // 哨兵位
// 	sentence = "-" + sentence

// 	for i := 1; i < len(sentence); i++ {
// 		for j := 0; j <= i; j++ {
// 			dp[i] = dp[i-1] + 1
// 			if find := dt.Find(sentence[j : i+1]); find {
// 				dp[i] = min(dp[j-1], dp[i])
// 				break
// 			}
// 		}
// 	}

// 	return dp[len(dp)-1]
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}

// 	return b
// }

// type Trie struct {
// 	Data      rune
// 	Children  []*Trie
// 	IsEndChar bool
// }

// func NewTrie() *Trie {
// 	return &Trie{
// 		Data:      '/',
// 		Children:  make([]*Trie, 26),
// 		IsEndChar: false,
// 	}
// }

// func (t *Trie) Insert(str string) {
// 	for _, v := range str {
// 		index := v - 'a'
// 		if t.Children[index] == nil {
// 			t.Children[index] = NewTrie()
// 			t.Children[index].Data = v
// 		}
// 		t = t.Children[index]
// 	}
// 	t.IsEndChar = true
// }

// func (t *Trie) Find(pattern string) bool {
// 	for _, v := range pattern {
// 		index := v - 'a'
// 		if v == '-' {
// 			return false
// 		}
// 		if t.Children[index] == nil {
// 			return false
// 		}
// 		t = t.Children[index]
// 	}
// 	return t.IsEndChar
// }

func respace(dictionary []string, sentence string) int {
	n, inf := len(sentence), 0x3f3f3f3f
	root := &Trie{next: [26]*Trie{}}
	for _, word := range dictionary {
		root.insert(word)
	}
	dp := make([]int, n+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = inf
	}
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		curPos := root
		for j := i; j >= 1; j-- {
			t := int(sentence[j-1] - 'a')
			if curPos.next[t] == nil {
				break
			} else if curPos.next[t].isEnd {
				dp[i] = min(dp[i], dp[j-1])
			}
			if dp[i] == 0 {
				break
			}
			curPos = curPos.next[t]
		}
	}
	return dp[n]
}

type Trie struct {
	next  [26]*Trie
	isEnd bool
}

func (this *Trie) insert(s string) {
	curPos := this
	for i := len(s) - 1; i >= 0; i-- {
		t := int(s[i] - 'a')
		if curPos.next[t] == nil {
			curPos.next[t] = &Trie{next: [26]*Trie{}}
		}
		curPos = curPos.next[t]
	}
	curPos.isEnd = true
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
