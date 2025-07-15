package sliding_window

import (
	"sync"

	"github.com/MordaTeam/go-toolbox/datastructs/dequeue"
)

type SlidingWindow[T any] interface {
	Push(el T)
	Clear()
	Data() []T
}

var _ SlidingWindow[any] = &slidingWindow[any]{}

type slidingWindow[T any] struct {
	mu        sync.RWMutex
	deq       dequeue.Dequeue[T]
	keepLastN uint
}

func New[T any](keepLastN uint) *slidingWindow[T] {
	return &slidingWindow[T]{
		keepLastN: keepLastN,
		deq:       dequeue.New[T](),
	}
}

func (w *slidingWindow[T]) Push(el T) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.deq.PushTail(el)
	if w.deq.Size() > w.keepLastN {
		_, _ = w.deq.PopHead()
	}
}

func (w *slidingWindow[T]) Clear() {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.deq.Clear()
}

func (w *slidingWindow[T]) Data() []T {
	w.mu.RLock()
	defer w.mu.RUnlock()

	return w.deq.GetAll()
}
