package comparer

// ComparerFunc lets create sateless comparer in short way.
type ComparerFunc[T any] func(a, b T) int

func (cf ComparerFunc[T]) Compare(a, b T) int {
	return cf(a, b)
}
