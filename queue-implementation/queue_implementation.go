package queueimplementation

import "fmt"

type QueueNode struct {
	value int
	next  string
}

type Queue struct {
	Length int
	head   QueueNode
	tail   QueueNode
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(value int) {
	fmt.Println("ENQ")
}

func (q *Queue) Dequeue() (int, bool) {
	fmt.Println("DEQ")
	return q.head.value, false
}

func (q *Queue) Peek() (int, bool) {
	if q.Length == 0 {
		return 0, false
	}
	return q.head.value, false
}
