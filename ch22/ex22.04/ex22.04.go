package main

import (
	"fmt"
)

type Stack struct {
	list []int
}

func (s *Stack) Push(value int) {
	s.list = append(s.list, value)
}

func (s *Stack) Pop() interface{} {
	if len(s.list) > 0 {
		lastIdx := len(s.list) - 1
		last := s.list[lastIdx]
		s.list = s.list[:lastIdx]
		return last
	}

	return nil
}

func NewStack() *Stack {
	return &Stack{}
}

func main() {
	stack := NewStack()

	for i := 1; i < 5; i++ {
		stack.Push(i)
	}

	value := stack.Pop()
	for value != nil {
		fmt.Printf("%v -> ", value)
		value = stack.Pop()
	}
}
