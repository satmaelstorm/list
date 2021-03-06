package list

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type stackSuite struct {
	suite.Suite
}

func TestStack(t *testing.T) {
	suite.Run(t, new(stackSuite))
}

func (s *stackSuite) TestStack() {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	s.Require().NotNil(stack.Top())
	s.Equal(3, stack.Top().Value())
	s.Equal(2, stack.Top().Next().Value())
	s.Equal(1, stack.Top().Next().Next().Value())
	s.Equal(3, stack.Len())

	r := stack.Pop()
	s.Require().NotNil(r)
	s.Equal(3, r.Value())
	s.Equal(2, stack.Top().Value())
	s.Equal(1, stack.Top().Next().Value())
	s.Equal(2, stack.Len())

	r = stack.Pop()
	s.Require().NotNil(r)
	s.Equal(2, r.Value())
	s.Equal(1, stack.Top().Value())
	s.Nil(stack.Top().Next())
	s.Equal(1, stack.Len())

	r = stack.Pop()
	s.Require().NotNil(r)
	s.Equal(1, r.Value())
	s.Nil(stack.Top())
	s.Equal(0, stack.Len())

	r = stack.Pop()
	s.Require().Nil(r)
	s.Equal(0, stack.Len())
}

func (s *stackSuite) TestRemove() {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	var n *Node[int]

	middle := stack.Top().Next().Next()
	s.Require().NotNil(middle)
	s.Equal(3, middle.Value())
	s.Equal(5, stack.Len())

	n = stack.Remove(middle)
	s.Equal(stack.Top().Next().Next(), n)
	s.Equal(5, stack.Top().Value())
	s.Equal(4, stack.Top().Next().Value())
	s.Equal(2, stack.Top().Next().Next().Value())
	s.Equal(1, stack.Top().Next().Next().Next().Value())
	s.Equal(4, stack.Len())

	n = stack.Remove(stack.Top().Next().Next().Next())
	s.Nil(n)
	s.Equal(5, stack.Top().Value())
	s.Equal(4, stack.Top().Next().Value())
	s.Equal(2, stack.Top().Next().Next().Value())
	s.Nil(stack.Top().Next().Next().Next())
	s.Equal(3, stack.Len())

	n = stack.Remove(stack.Top())
	s.Equal(stack.Top(), n)
	s.Equal(4, stack.Top().Value())
	s.Equal(2, stack.Top().Next().Value())
	s.Equal(2, stack.Len())

	n = stack.Remove(stack.Top())
	s.Equal(stack.Top(), n)
	s.Equal(2, stack.Top().Value())
	s.Equal(1, stack.Len())

	n = stack.Remove(stack.Top())
	s.Nil(n)
	s.Nil(stack.Top())
	s.Equal(0, stack.Len())

	n = stack.Remove(stack.Top())
	s.Nil(n)
	s.Equal(0, stack.Len())
}
