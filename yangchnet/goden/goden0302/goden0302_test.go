package goden0302

import (
	"fmt"
	"testing"
)

func Test_MinStack(t *testing.T) {
	s := Constructor()
	s.Push(3)
	s.Push(2)
	s.Push(1)
	fmt.Println(s.GetMin())
	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
	fmt.Println(s.GetMin())
}
