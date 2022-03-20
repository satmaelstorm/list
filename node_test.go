package list

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type nodeSuite struct {
	suite.Suite
}

func TestNode(t *testing.T) {
	suite.Run(t, new(nodeSuite))
}

func (s *nodeSuite) TestNodeValue() {
	node := new(Node[int])
	node2 := new(Node[int])
	node.SetValue(1)
	s.Equal(1, node.Value())
	s.Equal(0, node2.Value())
}

func (s *nodeSuite) TestNodeAddAfter() {
	head := new(Node[int])
	n := head.AddAfter(1)
	n = n.AddAfter(2)
	n.AddAfter(3)
	s.Equal(0, head.Value())
	s.Equal(1, head.Next().Value())
	s.Equal(2, head.Next().Next().Value())
	s.Equal(3, head.Next().Next().Next().Value())

	n.AddAfter(4)
	s.Equal(4, head.Next().Next().Next().Value())
	s.Equal(3, head.Next().Next().Next().Next().Value())
}

func (s *nodeSuite) TestNodeAddBefore() {
	head := new(Node[int])
	e := head.AddBefore(1)
	e1 := head.AddBefore(2)
	s.Equal(0, head.Value())
	s.Equal(2, head.Prev().Value())
	s.Equal(1, head.Prev().Prev().Value())

	e.AddBefore(3)

	s.Equal(2, head.Prev().Value())
	s.Equal(1, head.Prev().Prev().Value())
	s.Equal(3, head.Prev().Prev().Prev().Value())

	e1.AddBefore(4)

	s.Equal(2, head.Prev().Value())
	s.Equal(4, head.Prev().Prev().Value())
	s.Equal(1, head.Prev().Prev().Prev().Value())
	s.Equal(3, head.Prev().Prev().Prev().Prev().Value())
}

func (s *nodeSuite) TestNodeRemove() {
	head := new(Node[int])
	head.AddAfter(1).AddAfter(2).AddAfter(3).AddAfter(4)

	s.Equal(0, head.Value())
	s.Equal(1, head.Next().Value())
	s.Equal(2, head.Next().Next().Value())
	s.Equal(3, head.Next().Next().Next().Value())
	s.Equal(4, head.Next().Next().Next().Next().Value())

	el := head.Next().Next()
	s.Require().NotNil(el)

	p, n := el.Remove()
	s.Require().NotNil(p)
	s.Require().NotNil(n)
	s.Equal(1, p.Value())
	s.Equal(3, n.Value())

	s.Equal(0, head.Value())
	s.Equal(1, head.Next().Value())
	s.Equal(3, head.Next().Next().Value())
	s.Equal(4, head.Next().Next().Next().Value())

	el = head.Next().Next().Next()
	s.Require().NotNil(el)

	p, n = el.Remove()
	s.Require().NotNil(p)
	s.Require().Nil(n)
	s.Equal(3, p.Value())

	s.Equal(0, head.Value())
	s.Equal(1, head.Next().Value())
	s.Equal(3, head.Next().Next().Value())

	p, head = head.Remove()
	s.Require().Nil(p)
	s.Require().NotNil(head)
	s.Equal(1, head.Value())
	s.Equal(3, head.Next().Value())
}
