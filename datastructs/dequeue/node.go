package dequeue

type node[T any] struct {
	payload    T
	nextToTail *node[T]
	nextToHead *node[T]
}
