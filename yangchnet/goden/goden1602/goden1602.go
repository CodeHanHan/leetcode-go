package goden1602

type WordsFrequency struct {
	m map[string]int
}

func Constructor(book []string) WordsFrequency {
	_m := make(map[string]int)
	for i := 0; i < len(book); i++ {
		_m[book[i]]++
	}
	return WordsFrequency{
		m: _m,
	}
}

func (this *WordsFrequency) Get(word string) int {
	return this.m[word]
}

/**
 * Your WordsFrequency object will be instantiated and called as such:
 * obj := Constructor(book);
 * param_1 := obj.Get(word);
 */
