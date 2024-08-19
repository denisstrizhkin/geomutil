package geomutil

import "errors"

type BinTree[T any] struct {
	items []T
	comp  Comparator[T]
}

type Node[T any] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

func NewNode[T any](val T) *Node[T] {
	return &Node[T]{Value: val, Left: nil, Right: nil}
}

func (n *Node[T]) InsertLeftNode(val T) *Node[T] {
	node_l := NewNode(val)
	n.Left = node_l
	return node_l
}

func (n *Node[T]) InsertRightNode(val T) *Node[T] {
	node_r := NewNode(val)
	n.Right = node_r
	return node_r
}

func (n *Node[T]) IsLeafNode() bool {
	return n.Left == nil && n.Right == nil
}

func (n *Node[T]) IsFullNode() bool {
	return n.Left != nil && n.Right != nil
}

var ErrEmptyBinTree = errors.New("tree is empty")

type Comparator[T any] func(T, T) int

func NewBinTree[T any](comp Comparator[T]) *BinTree[T] {
	return &BinTree[T]{comp: comp}
}

func (s *BinTree[T]) Search(t T) (any, error) {
	l := len(s.items)
	if l == 0 {
		return nil, ErrEmptyBinTree
	}
	for i := 0; i < l; i++ {
		if s.comp(t, s.items[i]) == 0 {
			return t, nil
		} else if s.comp(t, s.items[i]) < 0 {

		}
	}
}

func ComparatorInt(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

func (s *BinTree[T]) Push(t T) {
	s.items = append(s.items, t)
}

func (s *BinTree[T]) Pop() (T, error) {
	l := len(s.items)
	var t T
	if l == 0 {
		return t, ErrEmptyBinTree
	}
	t = s.items[l-1]
	s.items = s.items[:l-1]
	return t, nil
}

func (s *BinTree[T]) Peek() (T, error) {
	l := len(s.items)
	var t T
	if l == 0 {
		return t, ErrEmptyBinTree
	}
	return s.items[l-1], nil
}

func (s *BinTree[T]) Length() int {
	return len(s.items)
}
