package queue

// Queue 队列
type Queue []interface{}

func NewQueue() *Queue {
	return new(Queue)
}

func (q *Queue) Add(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Pop() interface{} {
	data := (*q)[0]
	*q = (*q)[1:]
	return data
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Size() int {
	return len(*q)
}
