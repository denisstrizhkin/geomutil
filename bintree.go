package geomutil

import (
	"fmt"
)

type BinTree[Key, Value any] struct {
	head    *Node[Key, Value]
	comp    Comparator[Key]
	deleted Value
}

type Node[Key, Value any] struct {
	Key   Key
	Value Value
	Left  *Node[Key, Value]
	Right *Node[Key, Value]
}

func NewNode[Key, Value any](key Key, val Value) *Node[Key, Value] {
	return &Node[Key, Value]{Key: key, Value: val}
}

// func (n *Node[T]) insertLeftNode(val T) *Node[T] {
// 	n.Left = newNode(val)
// 	return n.Left
// }

// func (n *Node[T]) insertRightNode(val T) *Node[T] {
// 	n.Right = newNode(val)
// 	return n.Right
// }

func (n *Node[Key, Value]) IsLeafNode() bool {
	return n.Left == nil && n.Right == nil
}

func (n *Node[Key, Value]) IsFullNode() bool {
	return n.Left != nil && n.Right != nil
}

type Comparator[Key any] func(a, b Key) int

func NewBinTree[Key, Value any](comp Comparator[Key]) *BinTree[Key, Value] {
	return &BinTree[Key, Value]{comp: comp}
}

func (bt *BinTree[Key, Value]) InsertNode(key Key, val Value) {
	bt.head = bt.insertNode(bt.head, key, val)
}

func (bt *BinTree[Key, Value]) insertNode(
	node *Node[Key, Value], key Key, val Value,
) *Node[Key, Value] {
	if node == nil {
		fmt.Println("Inserted value!")
		return NewNode(key, val)
	}
	cmp := bt.comp(node.Key, key)
	switch {
	case cmp < 0:
		return bt.insertNode(node.Left, key, val)
	case cmp > 0:
		return bt.insertNode(node.Right, key, val)
	default:
		node.Value = val
		return node
	}
}

func (bt *BinTree[Key, Value]) Search(key Key) (Value, bool) {
	return bt.search(bt.head, key)
}

func (bt *BinTree[Key, Value]) search(node *Node[Key, Value], key Key) (Value, bool) {
	if node == nil {
		var zero Value
		return zero, false
	}
	cmp := bt.comp(node.Key, key)
	switch {
	case cmp < 0:
		return bt.search(node.Left, key)
	case cmp > 0:
		return bt.search(node.Right, key)
	default:
		return node.Value, true
	}
}

func (bt *BinTree[Key, Value]) DeleteNode(key Key) (Value, bool) {
	var zero Value
	var ok bool
	bt.deleted = zero
	bt.head, ok = bt.delete(bt.head, key)
	return bt.deleted, ok
}

func (bt *BinTree[Key, Value]) delete(node *Node[Key, Value], key Key) (*Node[Key, Value], bool) {
	if node == nil {
		return nil, false
	}
	cmp := bt.comp(node.Key, key)
	switch {
	case cmp < 0:
		return bt.delete(node.Left, key)
	case cmp > 0:
		return bt.delete(node.Right, key)
	case node.Left == nil:
		bt.deleted = node.Value
		return node.Right, true
	case node.Right == nil:
		bt.deleted = node.Value
		return node.Left, true
	default:
		// this is broken for now
		return node, true
	}
	// new_n := bt.minNodeSearch(node.Right)
	// bt.delete(node.Right, new_n.Value)
	// return new_n, nil
}

func (bt *BinTree[Key, Value]) minNodeSearch(node *Node[Key, Value]) *Node[Key, Value] {
	if node.Left == nil {
		return node
	}
	return bt.minNodeSearch(node.Left)
}
