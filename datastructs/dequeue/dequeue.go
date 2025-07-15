package dequeue

import (
	"errors"
)

var ErrDequeueIsEmpty = errors.New("dequeue is empty")

type Dequeue[T any] interface {
	Size() uint
	GetAll() []T
	PushHead(data T)
	PopHead() (T, error)
	GetHead() (T, error)
	PushTail(data T)
	PopTail() (T, error)
	GetTail() (T, error)
	Clear()
}

var _ Dequeue[any] = &dequeue[any]{}

type dequeue[T any] struct {
	head *node[T]
	tail *node[T]
	size uint
}

func New[T any]() *dequeue[T] {
	return &dequeue[T]{}
}

func (d *dequeue[T]) Size() uint {
	return d.size
}

func (d *dequeue[T]) GetAll() []T {
	res := make([]T, 0, d.size)
	for cur := d.head; cur != nil; cur = cur.nextToTail {
		res = append(res, cur.payload)
	}

	return res
}

func (d *dequeue[T]) Clear() {
	d.head = nil
	d.tail = nil
	d.size = 0
}

func (d *dequeue[T]) PushHead(data T) {
	defer func() {
		d.size++
	}()

	n := &node[T]{
		payload: data,
	}

	if d.size == 0 {
		d.head = n
		d.tail = n
		return
	}

	n.nextToTail = d.head
	d.head.nextToHead = n
	d.head = n
}

func (d *dequeue[T]) PopHead() (T, error) {
	if d.size == 0 {
		return *new(T), ErrDequeueIsEmpty
	}

	data := d.head.payload

	if d.size == 1 {
		d.head = nil
		d.tail = nil
		d.size--
		return data, nil
	}

	d.head.nextToTail.nextToHead = nil
	d.head = d.head.nextToTail
	d.size--
	return data, nil
}

func (d *dequeue[T]) GetHead() (T, error) {
	if d.size == 0 {
		return *new(T), ErrDequeueIsEmpty
	}

	return d.head.payload, nil
}

func (d *dequeue[T]) PushTail(data T) {
	defer func() {
		d.size++
	}()

	n := &node[T]{
		payload: data,
	}

	if d.size == 0 {
		d.head = n
		d.tail = n
		return
	}

	n.nextToHead = d.tail
	d.tail.nextToTail = n
	d.tail = n
}

func (d *dequeue[T]) PopTail() (T, error) {
	if d.size == 0 {
		return *new(T), ErrDequeueIsEmpty
	}

	data := d.head.payload

	if d.size == 1 {
		d.head = nil
		d.tail = nil
		d.size--
		return data, nil
	}

	d.tail.nextToHead.nextToTail = nil
	d.tail = d.tail.nextToHead
	d.size--
	return data, nil
}

func (d *dequeue[T]) GetTail() (T, error) {
	if d.size == 0 {
		return *new(T), ErrDequeueIsEmpty
	}

	return d.tail.payload, nil
}
