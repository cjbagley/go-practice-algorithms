package queueimplementation

type QueueNode struct {
	value int
	next  *QueueNode
}

type Queue struct {
	Length int
	head   *QueueNode
	tail   *QueueNode
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(value int) {
	q.Length++
	node := &QueueNode{
		value: value,
		next:  nil,
	}
	if q.head == nil {
		q.head = node
		q.tail = node
	}

	q.tail.next = node
	q.tail = node
}

func (q *Queue) Dequeue() (int, bool) {
	if q.head == nil {
		return 0, false
	}

	head := q.head
	q.head = head.next
	head.next = nil
	q.Length--

	return head.value, true
}

func (q *Queue) Peek() (int, bool) {
	if q.head == nil {
		return 0, false
	}
	return q.head.value, true
}
