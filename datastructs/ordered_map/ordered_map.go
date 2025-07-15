package ordered_map

import (
	"errors"
	"fmt"
	"sync"

	"github.com/MordaTeam/go-toolbox/datastructs/comparer"
	"github.com/MordaTeam/go-toolbox/datastructs/trees"
	"github.com/MordaTeam/go-toolbox/datastructs/trees/avl_tree"
)

type KeyNotFoundError[K any] = trees.NotFoundError[K]

type OrderedMap[K any, V any] interface {
	Set(key K, value V) error
	Get(key K) (V, error)
	Keys() []K
	Remove(key K) error
	Clear()
	Size() int
}

var _ OrderedMap[any, any] = &orderedMap[any, any]{}

type orderedMap[K any, V any] struct {
	mu   sync.RWMutex
	tree avl_tree.AVLTree[K, V]
}

func New[K, V any](comp comparer.Comparer[K]) *orderedMap[K, V] {
	return &orderedMap[K, V]{
		tree: avl_tree.NewWithCustomKey[K, V](comp),
	}
}

func (om *orderedMap[K, V]) Set(key K, value V) error {
	om.mu.Lock()
	defer om.mu.Unlock()

	if err := om.tree.Insert(key, value); err != nil {
		return fmt.Errorf("inserting to tree: %w", err)
	}
	return nil
}

func (om *orderedMap[K, V]) Get(key K) (V, error) {
	om.mu.RLock()
	defer om.mu.RUnlock()

	val, err := om.tree.Get(key)
	if errors.Is(err, trees.NotFoundError[K]{}) {
		return *new(V), KeyNotFoundError[K]{GivenKey: key}
	}
	if err != nil {
		return *new(V), fmt.Errorf("getting key from tree: %w", err)
	}

	return val, nil
}

func (om *orderedMap[K, V]) Keys() []K {
	om.mu.RLock()
	defer om.mu.RUnlock()
	return om.tree.Keys()
}

func (om *orderedMap[K, V]) Remove(key K) error {
	om.mu.Lock()
	defer om.mu.Unlock()

	err := om.tree.Remove(key)
	if errors.Is(err, trees.NotFoundError[K]{}) {
		return KeyNotFoundError[K]{GivenKey: key}
	}
	if err != nil {
		return fmt.Errorf("removing key from tree: %w", err)
	}

	return nil
}

func (om *orderedMap[K, V]) Clear() {
	om.mu.Lock()
	defer om.mu.Unlock()
	om.tree.Clear()
}

func (om *orderedMap[K, V]) Size() int {
	om.mu.RLock()
	defer om.mu.RUnlock()
	return int(om.tree.Size())
}
