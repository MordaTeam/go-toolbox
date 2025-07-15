package avl_tree

import (
	"errors"
	"fmt"
	"sync"

	"github.com/MordaTeam/go-toolbox/datastructs/comparer"
	"github.com/MordaTeam/go-toolbox/datastructs/trees"
)

type avlTree[K, V any] struct {
	mu sync.RWMutex

	comparer comparer.Comparer[K]

	root *node[K, V]
	size uint
}

// Creates new AVL tree with string key type.
func New[V any]() *avlTree[string, V] {
	return &avlTree[string, V]{
		comparer: comparer.NewStringComparer(),
	}
}

// Creates new AVL tree with custom key type.
// Corresponding Comparer is required.
func NewWithCustomKey[K, V any](comp comparer.Comparer[K]) *avlTree[K, V] {
	return &avlTree[K, V]{
		comparer: comp,
	}
}

// Adds new element to AVL tree.
// Fixes balance, if necessary.
func (t *avlTree[K, V]) Insert(key K, val V) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if err := t.insert(t.root, key, val); err != nil {
		return fmt.Errorf("inserting new element: %w", err)
	}

	t.size++
	return nil
}

// If tree contains given key, returns corresponding value.
// Returns ErrNotFound otherwise.
func (t *avlTree[K, V]) Get(key K) (V, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	n := t.get(key)
	if n == nil {
		return *new(V), trees.NotFoundError[K]{GivenKey: key}
	}

	return n.Value, nil
}

// If tree contains given key, removes KV pair from itself.
// Returns ErrNotFound otherwise.
func (t *avlTree[K, V]) Remove(key K) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	n := t.get(key)
	if n == nil {
		return trees.NotFoundError[K]{GivenKey: key}
	}
	t.remove(n)

	t.size--
	return nil
}

// Removes all keys from tree.
func (t *avlTree[K, V]) Clear() {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.root = nil
	t.size = 0
}

// Returns all keys that the tree contains in ascending order.
// Order is defined by comparer.
func (t *avlTree[K, V]) Keys() []K {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return t.keys(t.root)
}

// Returns number of nodes in tree.
func (t *avlTree[K, V]) Size() uint {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return t.size
}

// Return the height of tree root.
func (t *avlTree[K, V]) Height() uint {
	t.mu.RLock()
	defer t.mu.RUnlock()

	if t.root == nil {
		return 0
	}

	return t.root.height
}

func (t *avlTree[K, V]) insert(subroot *node[K, V], key K, val V) error {
	if subroot == nil {
		if subroot == t.root {
			t.root = newNode(key, val)
			return nil
		}
		return errors.New("given subroot is nil")
	}

	var nextNode *node[K, V]
	var nextNodeIsLeft bool
	switch t.comparer.Compare(key, subroot.Key) {
	case -1, 0:
		nextNode = subroot.rightChild
		nextNodeIsLeft = false
	case 1:
		nextNode = subroot.leftChild
		nextNodeIsLeft = true
	}

	if nextNode != nil {
		if err := t.insert(nextNode, key, val); err != nil {
			return err
		}
	} else {
		nextNode = newNode(key, val)
		nextNode.parent = subroot
		switch nextNodeIsLeft {
		case true:
			subroot.leftChild = nextNode
		case false:
			subroot.rightChild = nextNode
		}
	}

	t.balance(subroot)

	return nil
}

func (t *avlTree[K, V]) remove(subroot *node[K, V]) {
	if subroot == nil {
		return
	}

	if subroot.isLeaf() ||
		subroot.hasOnlyLeftChild() ||
		subroot.hasOnlyRightChild() {
		var child *node[K, V]
		if subroot.hasOnlyLeftChild() {
			child = subroot.leftChild
		}
		if subroot.hasOnlyRightChild() {
			child = subroot.rightChild
		}

		if subroot.parent == nil {
			t.root = child
			return
		}

		switch subroot {
		case subroot.parent.rightChild:
			subroot.parent.rightChild = child
		case subroot.parent.leftChild:
			subroot.parent.leftChild = child
		}

		if child != nil {
			child.parent = subroot.parent
		}
	} else {
		n := t.removeMin(subroot.rightChild)
		// keep all pointers on their place, just susbtitute the payload
		subroot.Key = n.Key
		subroot.Value = n.Value
	}

	for cur := subroot.parent; cur != nil; cur = cur.parent {
		t.balance(cur)
	}
}

func (t *avlTree[K, V]) removeMin(subroot *node[K, V]) *node[K, V] {
	if subroot.isLeaf() || subroot.hasOnlyRightChild() {
		// is the most left node
		t.remove(subroot)
		return subroot
	}

	n := t.removeMin(subroot.leftChild)
	t.balance(subroot)
	return n
}

func (t *avlTree[K, V]) get(key K) *node[K, V] {
	cur := t.root
	for cur != nil {
		switch t.comparer.Compare(key, cur.Key) {
		case -1:
			cur = cur.rightChild
		case 1:
			cur = cur.leftChild
		case 0:
			return cur
		}
	}
	return nil
}

func (t *avlTree[K, V]) keys(subroot *node[K, V]) []K {
	if subroot == nil {
		return []K{}
	}

	res := append(t.keys(subroot.leftChild), subroot.Key)
	res = append(res, t.keys(subroot.rightChild)...)
	return res
}

func (t *avlTree[K, V]) balance(subroot *node[K, V]) {
	subroot.fixHeight()
	switch subroot.balanceFactor() {
	case 2:
		if subroot.rightChild.balanceFactor() < 0 {
			// turns to greater left rotation
			t.rightRotation(subroot.rightChild)
		}
		t.leftRotation(subroot)
	case -2:
		if subroot.leftChild.balanceFactor() > 0 {
			// turns to greater right rotation
			t.leftRotation(subroot.leftChild)
		}
		t.rightRotation(subroot)
	default:
		return
	}
}

func (t *avlTree[K, V]) leftRotation(subroot *node[K, V]) {
	temp := subroot.rightChild
	if temp == nil {
		return
	}

	subroot.rightChild = temp.leftChild
	if subroot.rightChild != nil {
		subroot.rightChild.parent = subroot
	}

	temp.leftChild = subroot
	temp.parent = subroot.parent
	if subroot.parent != nil {
		switch subroot {
		case subroot.parent.leftChild:
			subroot.parent.leftChild = temp
		case subroot.parent.rightChild:
			subroot.parent.rightChild = temp
		}
	} else {
		t.root = temp
	}

	subroot.parent = temp

	subroot.fixHeight()
	temp.fixHeight()
}

func (t *avlTree[K, V]) rightRotation(subroot *node[K, V]) {
	temp := subroot.leftChild
	if temp == nil {
		return
	}

	subroot.leftChild = temp.rightChild
	if subroot.leftChild != nil {
		subroot.leftChild.parent = subroot
	}

	temp.rightChild = subroot
	temp.parent = subroot.parent
	if subroot.parent != nil {
		switch subroot {
		case subroot.parent.leftChild:
			subroot.parent.leftChild = temp
		case subroot.parent.rightChild:
			subroot.parent.rightChild = temp
		}
	} else {
		t.root = temp
	}

	subroot.parent = temp

	subroot.fixHeight()
	temp.fixHeight()
}
