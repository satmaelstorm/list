package list

// Node - double link list node
type Node[T any] struct {
	next, prev *Node[T]
	val        T
}

// Value - return value in node
func (n *Node[T]) Value() T {
	return n.val
}

// Next - return next node of list
func (n *Node[T]) Next() *Node[T] {
	return n.next
}

// Prev - return previous node of the list
func (n *Node[T]) Prev() *Node[T] {
	return n.prev
}

// SetValue - set value of node
func (n *Node[T]) SetValue(val T) {
	n.val = val
}

// AddAfter - add new node to the list after this node
func (n *Node[T]) AddAfter(nextVal T) *Node[T] {
	newNode := &Node[T]{val: nextVal, prev: n, next: n.next}
	n.next = newNode
	return newNode
}

// AddBefore - add new node to the list before this node
func (n *Node[T]) AddBefore(nextVal T) *Node[T] {
	newNode := &Node[T]{val: nextVal, next: n, prev: n.prev}
	n.prev = newNode
	return newNode
}

// Remove - remove node from the list
func (n *Node[T]) Remove() (prev, next *Node[T]) {
	prev = n.prev
	next = n.next

	if prev != nil {
		prev.next = next
	}

	if next != nil {
		next.prev = prev
	}

	return prev, next
}
