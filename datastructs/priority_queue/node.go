package priority_queue

type Node[T any] struct {
	Payload  T
	Priority uint
}
