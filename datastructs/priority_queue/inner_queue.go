package priority_queue

import (
	"container/heap"
)

type innerQueue[T any] []*Node[T]

var _ heap.Interface = &innerQueue[any]{}

// Implements heap.Interface
func (pq innerQueue[T]) Len() int {
	return len(pq)
}

// Implements heap.Interface
func (pq innerQueue[T]) Less(i, j int) bool {
	return pq[i].Priority > pq[j].Priority
}

// Implements heap.Interface
func (pq innerQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Implements heap.Interface
func (pq *innerQueue[T]) Push(data any) {
	//nolint: errcheck, forcetypeassert
	*pq = append(*pq, data.(*Node[T]))
}

// Implements heap.Interface
func (pq *innerQueue[T]) Pop() any {
	if len(*pq) == 0 {
		return nil
	}

	el := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return el
}
