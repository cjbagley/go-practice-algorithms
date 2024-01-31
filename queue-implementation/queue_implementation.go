package queueimplementation

import "fmt"

type QueueNode[T any] struct {
	value T
	next  string
}

type Queue[T any] struct {
	Length int
	head   QueueNode[T]
	tail   QueueNode[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(item T) {
	fmt.Println("ENQ")
}

func (q *Queue[T]) Dequeue() (T, bool) {
	fmt.Println("DEQ")
	return q.head.value, false
}

func (q *Queue[T]) Peek() (T, bool) {
	fmt.Println("Peek")
	return q.head.value, false
}
