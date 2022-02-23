package stack

import (
	"fmt"
	"testing"

	_ "github.com/spf13/pflag"
)

func Test_stack(t *testing.T) {
	s := NewStack()
	s.Push("1")
	fmt.Println(s)

	s.Push("2")
	s.Push("3")
	fmt.Println(s)

	s.Pop()
	fmt.Println(s)
	s.Pop()
	_, ok := s.Pop()
	if !ok {
		fmt.Println("stack is empty")
	}

	_, ok = s.Pop()
	if !ok {
		fmt.Println("stack is empty")
	}
}
