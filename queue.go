package list

// Queue - FIFO queue
type Queue[T any] struct {
	head, tail *Node[T]
	len        int
}

// NewQueue - create new queue
func NewQueue[T any]() *Queue[T] {
	return new(Queue[T])
}

func (q *Queue[T]) Head() *Node[T] {
	return q.head
}

func (q *Queue[T]) Tail() *Node[T] {
	return q.tail
}

func (q *Queue[T]) Enqueue(val T) {
	defer func() {
		q.len += 1
	}()
	if nil == q.head {
		q.head = &Node[T]{
			next: nil,
			prev: nil,
			val:  val,
		}
		q.tail = q.head
		return
	}
	q.tail = q.tail.AddAfter(val)
}

func (q *Queue[T]) Dequeue() *Node[T] {
	if nil == q.head {
		return nil
	}
	el := q.head
	q.head = el.Next()
	if nil == q.head {
		q.tail = nil
	}
	el.next = nil
	el.prev = nil
	q.len -= 1
	return el
}

func (q *Queue[T]) Remove(node *Node[T]) (prev, next *Node[T]) {
	if nil == node {
		return nil, nil
	}
	if q.head == node {
		q.head = node.next
	}
	if q.tail == node {
		q.tail = node.prev
	}
	q.len -= 1
	return node.Remove()
}

func (q *Queue[T]) MoveToBack(node *Node[T]) {
	if q.Tail() == node {
		return
	}
	q.Remove(node)
	q.Enqueue(node)
}

func (q *Queue[T]) Len() int {
	return q.len
}
