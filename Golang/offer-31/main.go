package main

import "math"

type stack struct {
	val []int
	top uint
}

func newStack() *stack {
	return &stack{
		val: []int{},
		top: 0,
	}
}

func (s *stack) push(val int) {
	s.val = append(s.val, val)
	s.top++
}

func (s *stack) isEmpty() bool {
	return s.top == 0
}

func (s *stack) pop() int {
	if s.top == 0 {
		panic("empty stack")
	}
	res := s.val[s.top-1]
	s.val = s.val[0 : len(s.val)-1]
	s.top--
	return res
}

func (s *stack) getTop() int {
	if s.isEmpty() {
		return int(math.NaN())
	}
	return s.val[s.top-1]
}

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 0 && len(popped) == 0 {
		return true
	}
	s := newStack()
	s.push(pushed[0])
	pushed = pushed[1:]

	for len(pushed) != 0 || s.getTop() == popped[0] {
		if s.getTop() == popped[0] { //目标值就在栈顶
			s.pop()
			popped = popped[1:]
			if len(popped) <= 0 {
				break
			}
			continue
		} else { //目标值不在栈顶
			s.push(pushed[0])
			if len(pushed) <= 1 {
				pushed = []int{}
			} else {
				pushed = pushed[1:]
			}

		}
	}
	if len(popped) <= 0 {
		return true
	}
	return false
}
