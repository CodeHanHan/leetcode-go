package trie

/**
@Author: yangchnet
@Create: 2022-03-08 13:44
@Description: Trie树，又叫字典树，常用于前缀匹配
*/

type Trie struct {
	Data      rune
	Children  []*Trie
	IsEndChar bool
}

func NewTrie() *Trie {
	return &Trie{
		Data:      '/',
		Children:  make([]*Trie, 26),
		IsEndChar: false,
	}
}

func (t *Trie) Insert(str string) {
	for _, v := range str {
		index := v - 'a'
		if t.Children[index] == nil {
			t.Children[index] = NewTrie()
			t.Children[index].Data = v
		}
		t = t.Children[index]
	}
	t.IsEndChar = true
}

func (t *Trie) Find(pattern string) bool {
	for _, v := range pattern {
		index := v - 'a'
		if t.Children[index] == nil {
			return false
		}
		t = t.Children[index]
	}
	return t.IsEndChar
}
