package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	list *list.List
}

func (q *Queue) Push(value int) {
	q.list.PushBack(value)
}

func (q *Queue) Pop() interface{} {
	front := q.list.Front()
	if front != nil {
		return q.list.Remove(front)
	}
	return nil
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func main() {
	queue := NewQueue()

	for i := 1; i < 5; i++ {
		queue.Push(i)
	}

	value := queue.Pop()
	for value != nil {
		fmt.Printf("%v -> ", value)
		value = queue.Pop()
	}
}
