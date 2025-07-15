package priority_queue

import (
	"container/heap"
	"errors"
	"fmt"
	"sync"

	"github.com/samber/lo"
)

var ErrNoElements = errors.New("queue is empty")

type PriorityQueue[T any] interface {
	Push(n Node[T])
	Pop() (Node[T], error)
	Size() int
	Items() []Node[T]
}

type priorityQueue[T any] struct {
	mu sync.RWMutex
	pq innerQueue[T]
}

func New[T any]() *priorityQueue[T] {
	pq := priorityQueue[T]{
		pq: make(innerQueue[T], 0), // TODO: initial cap from opt
	}

	heap.Init(&pq.pq)

	return &pq
}

func (pq *priorityQueue[T]) Push(n Node[T]) {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	heap.Push(&pq.pq, &n)
}

func (pq *priorityQueue[T]) Pop() (Node[T], error) {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	if len(pq.pq) == 0 {
		return Node[T]{}, ErrNoElements
	}

	n := heap.Pop(&pq.pq)

	res, ok := n.(*Node[T])
	if !ok {
		return Node[T]{}, fmt.Errorf("data has unexpected type: %T", n)
	}

	return *res, nil
}

func (pq *priorityQueue[T]) Size() int {
	pq.mu.RLock()
	defer pq.mu.RUnlock()
	return pq.pq.Len()
}

func (pq *priorityQueue[T]) Items() []Node[T] {
	pq.mu.RLock()
	defer pq.mu.RUnlock()
	return lo.Map(
		pq.pq,
		func(el *Node[T], _ int) Node[T] {
			return *el
		},
	)
}
