package stack

import "fmt"

type Stack struct {
	items []interface{}
	len   int
}

func NewStack() *Stack {
	return &Stack{
		items: make([]interface{}, 0),
		len:   0,
	}
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
	s.len++
}

func (s *Stack) Pop() (item interface{}, ok bool) {
	if s.Empty() {
		return nil, false
	}
	item = s.items[s.len-1]
	s.items = s.items[:s.len-1]
	s.len--
	return item, true
}

func (s *Stack) Top() (item interface{}, ok bool) {
	if s.Empty() {
		return nil, false
	}
	return s.items[0], true
}

func (s *Stack) Len() int {
	return s.len
}

func (s *Stack) Empty() bool {
	return s.len <= 0
}

func (s *Stack) String() string {
	return fmt.Sprint(s.items)
}
