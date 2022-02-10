package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	list *list.List
}

func (s *Stack) Push(value int) {
	s.list.PushBack(value)
}

func (s *Stack) Pop() interface{} {
	back := s.list.Back()
	if back != nil {
		return s.list.Remove(back)
	}
	return nil
}

func NewStack() *Stack {
	return &Stack{list.New()}
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
