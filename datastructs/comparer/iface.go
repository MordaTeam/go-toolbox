package comparer

// Comparer provides a way to compare 2 elements.
type Comparer[T any] interface {
	// Returns
	// -1 if a > b
	// 0 if a = b
	// 1 if a < b
	Compare(a, b T) int
}
