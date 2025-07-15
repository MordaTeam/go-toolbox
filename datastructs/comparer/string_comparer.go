package comparer

var _ Comparer[string] = &stringComparer{}

type stringComparer struct{}

// Creates new comparer that works withs strings.
func NewStringComparer() *stringComparer {
	return &stringComparer{}
}

func (comp *stringComparer) Compare(a, b string) int {
	if a > b {
		return -1
	}
	if a < b {
		return 1
	}
	return 0
}
