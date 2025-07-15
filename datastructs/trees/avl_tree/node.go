package avl_tree

type node[K, V any] struct {
	Key   K
	Value V

	leftChild  *node[K, V]
	rightChild *node[K, V]
	parent     *node[K, V]

	height uint
}

func newNode[K, V any](key K, val V) *node[K, V] {
	return &node[K, V]{
		Key:    key,
		Value:  val,
		height: 1,
	}
}

func (n *node[K, V]) isLeaf() bool {
	return n.leftChild == nil && n.rightChild == nil
}

func (n *node[K, V]) hasOnlyLeftChild() bool {
	return n.leftChild != nil && n.rightChild == nil
}

func (n *node[K, V]) hasOnlyRightChild() bool {
	return n.leftChild == nil && n.rightChild != nil
}

func (n *node[K, V]) balanceFactor() int {
	l, r := 0, 0
	if n.leftChild != nil {
		l = int(n.leftChild.height)
	}
	if n.rightChild != nil {
		r = int(n.rightChild.height)
	}
	return r - l
}

func (n *node[K, V]) fixHeight() {
	if n.isLeaf() {
		n.height = 1
		return
	}

	var l, r uint = 0, 0
	if n.leftChild != nil {
		l = n.leftChild.height
	}
	if n.rightChild != nil {
		r = n.rightChild.height
	}

	n.height = max(l, r) + 1
}
