package ds

type Queue struct {
	elements []interface{}
}

func (q *Queue) Size() int {
	return len(q.elements)
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue) Peek() interface{} {
	if q.Size() == 0 {
		return nil
	}
	return q.elements[0]
}

func (q *Queue) Enqueue(vars ...interface{}) {
	q.elements = append(q.elements, vars...)
}

func (q *Queue) Dequeue() interface{} {
	if q.Size() == 0 {
		return nil
	}

	val := q.elements[0]
	if q.Size() == 1 {
		q.elements = nil
	} else {
		q.elements = q.elements[1:]
	}
	return val
}
