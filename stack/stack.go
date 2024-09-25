package stack

type entry[E comparable] struct {
	next  *entry[E]
	value E
}

type Stack[E comparable] struct {
	top    *entry[E]
	length int
}

func NewStack[E comparable]() *Stack[E] {
	return &Stack[E]{}
}

func (s *Stack[E]) Empty() bool {
	return s.top == nil
}

func (s *Stack[E]) IndexOf(e E) int {
	i := 0
	for entry := s.top; entry != nil; entry = entry.next {
		if entry.value == e {
			return i
		}
		i++
	}
	return -1
}

func (s *Stack[E]) Len() int {
	return s.length
}

func (s *Stack[E]) Push(e E) {
	s.top = &entry[E]{
		next:  s.top,
		value: e,
	}
	s.length++
}

func (s *Stack[E]) Pop() (E, bool) {
	if s.top == nil {
		var e E
		return e, false
	}

	top := s.top
	s.top = top.next

	top.next = nil
	s.length--

	return top.value, true
}

func (s *Stack[E]) Peek() (E, bool) {
	if s.top == nil {
		var e E
		return e, false
	}

	top := s.top
	return top.value, true
}
