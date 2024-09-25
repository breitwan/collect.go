package queue

type entry[E comparable] struct {
	next  *entry[E]
	value E
}

type Queue[E comparable] struct {
	head, tail *entry[E]
	length     int
}

func NewQueue[E comparable]() *Queue[E] {
	return &Queue[E]{}
}

func (q *Queue[E]) Clear() {
	q.head, q.tail = nil, nil
	q.length = 0
}

func (q *Queue[E]) Len() int {
	return q.length
}

func (q *Queue[E]) Peek() (E, bool) {
	entry := q.head
	if entry == nil {
		var e E
		return e, false
	}

	return entry.value, true
}

func (q *Queue[E]) Push(e E) {
	entry := &entry[E]{value: e}

	q.length++

	if q.head == nil {
		q.head = entry
		q.tail = q.head
		return
	}

	q.tail.next = entry
	q.tail = q.tail.next
}

func (q *Queue[E]) Pop() (E, bool) {
	entry := q.head
	if entry == nil {
		var e E
		return e, false
	}

	next := q.head.next
	q.head = next
	q.length--

	return entry.value, true
}
