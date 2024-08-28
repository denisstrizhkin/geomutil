package geomutil

import (
	"errors"
	"fmt"
)

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

// func (n *Node[T]) insertLeftNode(val T) *Node[T] {
// 	n.Left = newNode(val)
// 	return n.Left
// }

// func (n *Node[T]) insertRightNode(val T) *Node[T] {
// 	n.Right = newNode(val)
// 	return n.Right
// }

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

func (bt *BinTree[T]) insertNode(node *Node[T], val T) *Node[T] {
	if node == nil {
		fmt.Println("Inserted value!")
		return newNode(val)
	}
	if bt.comp(val, node.Value) < 0 {
		fmt.Println("Went left...")
		return bt.insertNode(node.Left, val)
	}
	if bt.comp(val, node.Value) > 0 {
		fmt.Println("Went right...")
		return bt.insertNode(node.Right, val)
	}
	return node
}

func (bt *BinTree[T]) InsertNode(val T) {
	bt.head = bt.insertNode(bt.head, val)
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

func (bt *BinTree[T]) delete(node *Node[T], val T) (any, error) {
	if node == nil {
		return nil, ErrItemNotFound
	}
	if bt.comp(val, node.Value) > 0 {
		return bt.delete(node.Right, val)
	} else if bt.comp(val, node.Value) < 0 {
		return bt.delete(bt.head.Left, val)
	}
	if node.Left == nil {
		return node.Right, nil
	} else if node.Right == nil {
		return node.Right, nil
	}
	new_n := bt.minNodeSearch(node.Right)
	bt.delete(node.Right, new_n.Value)
	return new_n, nil
}

func (bt *BinTree[T]) DeleteNode(val T) (any, error) {
	return bt.delete(bt.head, val)
}

func (bt *BinTree[T]) minNodeSearch(node *Node[T]) *Node[T] {
	if node.Left == nil {
		return node
	}
	return bt.minNodeSearch(node.Left)
}
