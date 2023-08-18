package queue

type Queue struct {
	items []interface{}
}

func (q *Queue) Push(item interface{}) interface{} {
	q.items = append(q.items, item)
	return item
}

func (q *Queue) Pop() interface{} {
	if q.Len() > 0 {
		item := q.items[0]
		q.items = q.items[1:]
		return item
	}
	return nil
}

func (q *Queue) Peek() interface{} {
	if q.Len() > 0 {
		return q.items[0]
	}
	return nil
}

func (q *Queue) Len() int {
	return len(q.items)
}
