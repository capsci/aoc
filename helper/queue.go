package helper

type Queue []interface{}

func (q *Queue) Enqueue(val interface{}) {
	*q = append(*q, val)
}

func (q *Queue) Dequeue() (val interface{}) {
	val = (*q)[0]
	*q = (*q)[1:]
	return val
}

func (q *Queue) Top() (val interface{}) {
	return (*q)[0]
}

func (q *Queue) Empty() bool {
	return q.Size() == 0
}

func (q *Queue) Size() int {
	return len(*q)
}
