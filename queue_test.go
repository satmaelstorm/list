package list

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type queueSuite struct {
	suite.Suite
}

func TestQueue(t *testing.T) {
	suite.Run(t, new(queueSuite))
}

func (s *queueSuite) TestQueue() {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	s.Equal(1, q.Head().Value())
	s.Equal(3, q.Tail().Value())
	s.Equal(3, q.Len())

	el := q.Dequeue()
	s.Require().NotNil(el)
	s.Equal(1, el.Value())
	s.Nil(el.Next())
	s.Nil(el.Prev())
	s.Equal(2, q.Head().Value())
	s.Equal(3, q.Tail().Value())
	s.Equal(2, q.Len())

	el = q.Dequeue()
	s.Require().NotNil(el)
	s.Equal(2, el.Value())
	s.Equal(3, q.Head().Value())
	s.Equal(3, q.Tail().Value())
	s.Equal(1, q.Len())

	el = q.Dequeue()
	s.Require().NotNil(el)
	s.Equal(3, el.Value())
	s.Nil(q.Head())
	s.Nil(q.Tail())
	s.Equal(0, q.Len())

	el = q.Dequeue()
	s.Nil(el)
	s.Equal(0, q.Len())
}

func (s *queueSuite) TestQueueRemove() {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	var w *Node[int]
	p, n := q.Remove(w)
	s.Nil(p)
	s.Nil(n)

	s.Equal(1, q.Head().Value())
	s.Equal(2, q.Head().Next().Value())
	s.Equal(3, q.Head().Next().Next().Value())
	s.Equal(4, q.Head().Next().Next().Next().Value())
	s.Equal(5, q.Head().Next().Next().Next().Next().Value())
	s.NotNil(q.Head())
	s.NotNil(q.Tail())
	s.Equal(5, q.Len())

	q.Remove(q.Head().Next().Next())
	s.Equal(1, q.Head().Value())
	s.Equal(2, q.Head().Next().Value())
	s.Equal(4, q.Head().Next().Next().Value())
	s.Equal(5, q.Head().Next().Next().Next().Value())
	s.NotNil(q.Head())
	s.NotNil(q.Tail())
	s.Equal(4, q.Len())

	q.Remove(q.Head())
	s.Equal(2, q.Head().Value())
	s.Equal(4, q.Head().Next().Value())
	s.Equal(5, q.Head().Next().Next().Value())
	s.NotNil(q.Head())
	s.NotNil(q.Tail())
	s.Equal(3, q.Len())

	q.Remove(q.Tail())
	s.Equal(2, q.Head().Value())
	s.Equal(4, q.Head().Next().Value())
	s.NotNil(q.Head())
	s.NotNil(q.Tail())
	s.Equal(2, q.Len())

	q.Remove(q.Head().Next())
	s.Equal(2, q.Head().Value())
	s.NotNil(q.Head())
	s.NotNil(q.Tail())
	s.Equal(1, q.Len())

	q.Remove(q.Head())
	s.Nil(q.Head())
	s.Nil(q.Tail())
	s.Equal(0, q.Len())
}

func (s *queueSuite) TestMoveToBack() {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	e := q.Head().Next().Next()
	s.Equal(3, e.Value())

	q.MoveToBack(e)

	s.Equal(1, q.Head().Value())
	s.Equal(2, q.Head().Next().Value())
	s.Equal(4, q.Head().Next().Next().Value())
	s.Equal(5, q.Head().Next().Next().Next().Value())
	s.Equal(3, q.Head().Next().Next().Next().Next().Value())
	s.NotNil(q.Head())
	s.NotNil(q.Tail())
	s.Equal(5, q.Len())

	q.MoveToBack(q.Head())

	s.Equal(2, q.Head().Value())
	s.Equal(4, q.Head().Next().Value())
	s.Equal(5, q.Head().Next().Next().Value())
	s.Equal(3, q.Head().Next().Next().Next().Value())
	s.Equal(1, q.Head().Next().Next().Next().Next().Value())
	s.NotNil(q.Head())
	s.NotNil(q.Tail())
	s.Equal(5, q.Len())

	q.MoveToBack(q.Tail())

	s.Equal(2, q.Head().Value())
	s.Equal(4, q.Head().Next().Value())
	s.Equal(5, q.Head().Next().Next().Value())
	s.Equal(3, q.Head().Next().Next().Next().Value())
	s.Equal(1, q.Head().Next().Next().Next().Next().Value())
	s.NotNil(q.Head())
	s.NotNil(q.Tail())
	s.Equal(5, q.Len())
}
