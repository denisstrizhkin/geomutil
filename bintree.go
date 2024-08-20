package geomutil

import "errors"

type BinTree[T any] struct {
	head *Node[T]
	comp Comparator[T]
}

type Node[T any] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

func newNode[T any](val T) *Node[T] {
	return &Node[T]{Value: val}
}

func (n *Node[T]) InsertLeftNode(val T) *Node[T] {
	n.Left = newNode(val)
	return n.Left
}

func (n *Node[T]) InsertRightNode(val T) *Node[T] {
	n.Right = newNode(val)
	return n.Right
}

func (n *Node[T]) IsLeafNode() bool {
	return n.Left == nil && n.Right == nil
}

func (n *Node[T]) IsFullNode() bool {
	return n.Left != nil && n.Right != nil
}

var ErrEmptyBinTree = errors.New("bin tree is empty")
var ErrItemNotFound = errors.New("no such item in your bin tree")

type Comparator[T any] func(a, b T) int

func NewBinTree[T any](comp Comparator[T]) *BinTree[T] {
	return &BinTree[T]{comp: comp}
}

func (bt *BinTree[T]) search(node *Node[T], val T) (any, error) {
	if node == nil {
		return nil, ErrItemNotFound
	}
	if bt.comp(val, node.Value) == 0 {
		return val, nil
	}
	if bt.comp(val, node.Value) < 0 {
		return bt.search(node.Left, val)
	}
	return bt.search(node.Right, val)
}

func (bt *BinTree[T]) Search(val T) (any, error) {
	return bt.search(bt.head, val)
}
