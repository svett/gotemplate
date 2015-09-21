package queue 

// template type Queue(A)
type A int

type QueueElement struct {
	next *QueueElement
	Value A
}

type Queue struct {
	head *QueueElement
}

func (q *Queue) Head() *QueueElement {
	return q.head
}

