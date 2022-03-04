package goden0301

import (
	"testing"
)

func Test_Triple(t *testing.T) {
	tt := Constructor(2)
	tt.Push(0, 1)
	tt.Push(0, 2)
	tt.Push(0, 3)
	tt.Push(1, 4)
	tt.Push(1, 5)
	tt.Push(1, 5)
	tt.Push(2, 6)
	tt.Push(2, 7)
	tt.Push(2, 8)
	tt.Pop(0)
	tt.Pop(0)
	tt.Pop(0)
	tt.Pop(1)
	tt.Pop(1)
	tt.Pop(1)
	tt.Pop(2)
	tt.Pop(2)
	tt.Pop(2)
	tt.IsEmpty(0)

}
