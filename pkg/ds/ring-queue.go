package ds

type queue interface {
	push(interface{})
	pop() interface{}
	list() []interface{}
}

func NewRingQueue(size int) *ringQueue {
	if size <= 0 {
		return nil
	}

	return &ringQueue{
		size: size,
		ring: make([]interface{}, size),
	}
}

type ringQueue struct {
	size       int
	head, tail int
	ring       []interface{}
	full       bool
}

func (q *ringQueue) isFull() bool {
	return q.full && q.head == q.tail
}

func (q *ringQueue) isEmpty() bool {
	return !q.full && q.head == q.tail
}

func (q *ringQueue) push(node interface{}) {
	if q.isFull() {
		return
	}

	q.ring[q.tail] = node

	q.tail = (q.tail + 1) % q.size

	if q.tail == q.head {
		q.full = true
	}
}

func (q *ringQueue) pop() interface{} {
	if q.isEmpty() {
		return nil
	}

	if q.tail == q.head {
		q.full = false
	}

	res := q.ring[q.head]
	q.ring[q.head] = nil
	q.head = (q.head + 1) % q.size

	return res
}

func (q *ringQueue) list() []interface{} {
	res := make([]interface{}, 0)

	if q.isEmpty() {
		return res
	}

	for i := 0; (q.isFull() && i < q.size) ||
		(!q.isFull() && (q.head+i)%q.size != q.tail); i++ {
		res = append(res, q.ring[(q.head+i)%q.size])
	}

	return res
}
