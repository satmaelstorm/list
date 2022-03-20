package list

type Stack[T any] struct {
	top *Node[T]
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
}

func (s *Stack[T]) Pop() *Node[T] {
	if nil == s.top {
		return nil
	}
	r := s.top
	s.top = r.Next()
	r.next = nil
	r.prev = nil
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
	return n
}
