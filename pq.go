package list

import (
	"errors"
	"golang.org/x/exp/constraints"
)

var ErrPqIsFull = errors.New("PQ is full")

type PqItem[K constraints.Integer, T any] struct {
	orderBy K
	value   T
	pos     int
}

func (p *PqItem[K, T]) GetIndex() int {
	return p.pos
}

func (p *PqItem[K, T]) GetOrderBy() K {
	return p.orderBy
}

func (p *PqItem[K, T]) GetValue() T {
	return p.value
}

func (p *PqItem[K, T]) SetValue(value T) {
	p.value = value
}

type PQ[K constraints.Integer, T any] struct {
	items         []*PqItem[K, T]
	currentLength int
	capacity      int
}

func NewPQ[K constraints.Integer, T any](maxLength int) *PQ[K, T] {
	return &PQ[K, T]{
		items:    make([]*PqItem[K, T], maxLength+1),
		capacity: maxLength,
	}
}

func (p *PQ[K, T]) less(i, j int) bool {
	return p.items[i].orderBy < p.items[j].orderBy
}

func (p *PQ[K, T]) exch(i, j int) {
	p.items[i], p.items[j] = p.items[j], p.items[i]
	p.items[i].pos = i
	p.items[j].pos = j
}

func (p *PQ[K, T]) swim(k int) {
	for k > 1 && p.less(k/2, k) {
		p.exch(k/2, k)
		k = k / 2
	}
}

func (p *PQ[K, T]) sink(k int) {
	for 2*k <= p.currentLength {
		j := 2 * k
		if j < p.currentLength && p.less(j, j+1) {
			j++
		}
		if !p.less(k, j) {
			break
		}
		p.exch(k, j)
		k = j
	}
}

func (p *PQ[K, T]) insert(orderBy K, value T) *PqItem[K, T] {
	p.currentLength += 1
	item := &PqItem[K, T]{orderBy: orderBy, value: value, pos: p.currentLength}
	p.items[p.currentLength] = item
	p.swim(p.currentLength)
	return item
}

func (p *PQ[K, T]) Enqueue(orderBy K, value T) (*PqItem[K, T], error) {
	if p.currentLength >= p.capacity {
		return nil, ErrPqIsFull
	}
	return p.insert(orderBy, value), nil
}

func (p *PQ[K, T]) EnqueueWithOust(orderBy K, value T) (*PqItem[K, T], *PqItem[K, T]) {
	var oust *PqItem[K, T]
	if p.currentLength >= p.capacity {
		oust, _ = p.Dequeue()
	}
	return p.insert(orderBy, value), oust
}

func (p *PQ[K, T]) Dequeue() (*PqItem[K, T], bool) {
	if p.currentLength < 1 {
		return nil, false
	}
	result := p.items[1]
	p.Delete(result)
	return result, true
}

func (p *PQ[K, T]) Delete(item *PqItem[K, T]) {
	pos := item.pos
	p.exch(item.pos, p.currentLength)
	p.currentLength -= 1
	p.items[p.currentLength+1] = nil
	p.sink(pos)
}

func (p *PQ[K, T]) CurrentLength() int {
	return p.currentLength
}

func (p *PQ[K, T]) Cap() int {
	return p.capacity
}

func (p *PQ[K, T]) IncInPosition(k int) *PqItem[K, T] {
	p.items[k].orderBy += 1
	res := p.items[k]
	p.swim(k)
	return res
}

func (p *PQ[K, T]) DecInPosition(k int) *PqItem[K, T] {
	p.items[k].orderBy -= 1
	res := p.items[k]
	p.sink(k)
	return res
}
