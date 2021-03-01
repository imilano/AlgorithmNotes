package utils

type Queue struct {
	que []interface{}
}

func (q *Queue) Enqueue(v interface{}) {
	q.que = append(q.que, v)
}

func (q *Queue) Dequeue() interface{} {
	var res interface{}
	if !q.IsEmpty() {
		res = q.que[0]
		q.que = q.que[1:]
	}

	return res
}

func (q *Queue) IsEmpty() bool {
	return len(q.que) == 0
}
