package list

import (
	"github.com/stretchr/testify/suite"
	"math/rand"
	"testing"
	"time"
)

type pqSuite struct {
	suite.Suite
}

func TestPq(t *testing.T) {
	suite.Run(t, new(pqSuite))
}

func (s *pqSuite) SetupSuite() {
	rand.Seed(time.Now().Unix())
}

func (s *pqSuite) TestPqFill() {
	pq := NewPQ[int, int](5)
	for i := 0; i < 5; i++ {
		_, err := pq.Enqueue(rand.Intn(10), i)
		s.NoError(err)
	}
	_, err := pq.Enqueue(rand.Intn(10), 6)
	s.Error(err)

	_, r0 := pq.EnqueueWithOust(rand.Intn(10), 7)
	s.Require().NotNil(r0)

	r1, ok := pq.Dequeue()
	s.Require().True(ok)
	r2, ok := pq.Dequeue()
	s.Require().True(ok)
	s.True(r1.GetOrderBy() >= r2.GetOrderBy())
	s.True(r0.GetOrderBy() >= r2.GetOrderBy())
	r3, ok := pq.Dequeue()
	s.Require().True(ok)
	s.True(r2.GetOrderBy() >= r3.GetOrderBy())
	r4, ok := pq.Dequeue()
	s.Require().True(ok)
	s.True(r3.GetOrderBy() >= r4.GetOrderBy())
	r5, ok := pq.Dequeue()
	s.Require().True(ok)
	s.True(r4.GetOrderBy() >= r5.GetOrderBy())
	r6, ok := pq.Dequeue()
	s.Require().False(ok)
	s.Nil(r6)
}

func (s *pqSuite) TestPqInc() {
	pq := NewPQ[int, int](5)
	_, _ = pq.Enqueue(2, 1)
	_, _ = pq.Enqueue(2, 1)
	_, _ = pq.Enqueue(2, 1)
	_, _ = pq.Enqueue(2, 1)
	_, _ = pq.Enqueue(1, 9)
	s.Equal(1, pq.items[1].value)
	s.Equal(9, pq.items[5].value)
	e := pq.IncInPosition(5)
	s.Equal(9, e.value)
	s.Equal(2, e.orderBy)
	pq.IncInPosition(e.pos)
	s.Equal(3, e.orderBy)
	s.Equal(1, e.pos)
}

func (s *pqSuite) TestPqDec() {
	pq := NewPQ[int, int](5)
	_, _ = pq.Enqueue(2, 1)
	_, _ = pq.Enqueue(2, 1)
	_, _ = pq.Enqueue(2, 1)
	_, _ = pq.Enqueue(2, 1)
	_, _ = pq.Enqueue(3, 9)
	s.Equal(9, pq.items[1].value)
	s.Equal(1, pq.items[5].value)
	e := pq.DecInPosition(1)
	s.Equal(9, e.value)
	s.Equal(2, e.orderBy)
	pq.DecInPosition(e.pos)
	s.Equal(1, e.orderBy)
	s.True(e.pos > 1)
}
