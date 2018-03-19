package stack

// Stack ...
type Stack struct {
	container []interface{}
	len       int
}

// New ...
func New() *Stack {
	return &Stack{
		container: make([]interface{}, 0),
	}
}

// Each ...
func (s *Stack) Each(f func(interface{}) bool) {
	for i := s.len - 1; i >= 0; i-- {
		if !f(s.container[i]) {
			return
		}
	}
}

// Iterator ...
type Iterator struct {
	idx   int
	stack *Stack
}

// Iterator ...
func (s *Stack) Iterator() *Iterator {
	return &Iterator{
		idx:   s.Len() - 1,
		stack: s,
	}
}

// Next ...
func (i *Iterator) Next() interface{} {
	i.idx--
	if i.idx < 0 {
		return nil
	}
	return i.stack.container[i.idx]
}

// Elems ...
func (s *Stack) Elems() <-chan interface{} {
	e := make(chan interface{})
	go func() {
		for i := s.len - 1; i >= 0; i-- {
			e <- s.container[i]
		}
		close(e)
	}()
	return e
}

// Fall ...
func (s *Stack) Fall(n int) {
	s.len -= n
	if s.len < 0 {
		s.len = 0
	}
}

// Rise ...
func (s *Stack) Rise(n int) {
	s.len += n
	if len(s.container) < s.len {
		s.len = len(s.container)
	}
}

// PopN ...
func (s *Stack) PopN(n int) []interface{} {
	if s.len == 0 {
		return nil
	}
	s.len -= n
	elems := s.container[s.len:]
	s.container = s.container[:s.len]
	return elems
}

// Push ...
func (s *Stack) Push(v interface{}) {
	s.container = append(s.container, v)
	s.len++
}

// Pop ...
func (s *Stack) Pop() interface{} {
	s.len--
	return s.container[s.len]
}

// Peak ...
func (s *Stack) Peak() interface{} {
	return s.container[s.len-1]
}

// Len ...
func (s *Stack) Len() int {
	return s.len
}

// Empty ...
func (s *Stack) Empty() bool {
	return s.len == 0
}
