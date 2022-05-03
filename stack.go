package list

type Stack[T any] struct {
	top *Node[T]
	len int
}

func NewStack[T any]() *Stack[T] {
	return new(Stack[T])
}

func (s *Stack[T]) Top() *Node[T] {
	return s.top
}

func (s *Stack[T]) Push(val T) {
	node := &Node[T]{val: val, prev: nil, next: s.top}
	if nil != s.top {
		s.top.prev = node
	}
	s.top = node
	s.len += 1
}

func (s *Stack[T]) Pop() *Node[T] {
	if nil == s.top {
		return nil
	}
	r := s.top
	s.top = r.Next()
	r.next = nil
	r.prev = nil
	s.len -= 1
	return r
}

func (s *Stack[T]) Remove(node *Node[T]) (next *Node[T]) {
	if nil == node {
		return nil
	}
	if s.top == node {
		s.top = node.Next()
	}
	_, n := node.Remove()
	s.len -= 1
	return n
}

func (s *Stack[T]) Len() int {
	return s.len
}
