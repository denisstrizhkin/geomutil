package bintree

import (
	"fmt"
	"strings"
)

type Node[Key, Value any] struct {
	Key    Key
	Value  Value
	Left   *Node[Key, Value]
	Right  *Node[Key, Value]
	height int
}

func NewNode[Key, Value any](key Key, val Value) *Node[Key, Value] {
	return &Node[Key, Value]{height: 1, Key: key, Value: val}
}

// func (n *Node[T]) insertLeftNode(val T) *Node[T] {
// 	n.Left = newNode(val)
// 	return n.Left
// }

// func (n *Node[T]) insertRightNode(val T) *Node[T] {
// 	n.Right = newNode(val)
// 	return n.Right
// }

func (n *Node[Key, Value]) String() string {
	return fmt.Sprintf("(%v| %v: %v)", n.height, n.Key, n.Value)
}

type Comparator[Key any] func(a, b Key) int

type BinTree[Key, Value any] struct {
	head  *Node[Key, Value]
	cmp   Comparator[Key]
	count int
}

func NewBinTree[Key, Value any](comp Comparator[Key]) *BinTree[Key, Value] {
	return &BinTree[Key, Value]{cmp: comp}
}

func (bt *BinTree[Key, Value]) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprint(bt.Size(), "| [ "))
	bt.PreOrderTraversalNode(func(node *Node[Key, Value]) {
		sb.WriteString(node.String())
		sb.WriteRune(' ')
	})
	sb.WriteRune(']')
	return sb.String()
}

func (n *Node[Key, Value]) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}

func (n *Node[Key, Value]) IsFull() bool {
	return n.Left != nil && n.Right != nil
}

func (n *Node[Key, Value]) minNode() *Node[Key, Value] {
	if n.Left == nil {
		return n
	}
	return n.minNode()
}

func (bt *BinTree[Key, Value]) Size() int {
	return bt.count
}

func (n *Node[Key, Value]) leftRotation() *Node[Key, Value] {
	y := n.Right
	n.Right = y.Left
	y.Left = n
	n.updateHeight()
	y.updateHeight()
	return y
}

func (n *Node[Key, Value]) rightRotation() *Node[Key, Value] {
	y := n.Left
	n.Left = y.Right
	y.Right = n
	n.updateHeight()
	y.updateHeight()
	return y
}

func (bt *BinTree[Key, Value]) Put(key Key, val Value) {
	bt.head = bt.put(bt.head, key, val)
}

func (bt *BinTree[Key, Value]) put(
	node *Node[Key, Value], key Key, val Value,
) *Node[Key, Value] {
	if node == nil {
		bt.count++
		return NewNode(key, val)
	}
	cmp := bt.cmp(node.Key, key)
	switch {
	case cmp > 0:
		node.Left = bt.put(node.Left, key, val)
	case cmp < 0:
		node.Right = bt.put(node.Right, key, val)
	default:
		node.Value = val
	}
	if absInt(node.getBF()) > 1 {
		node = bt.balance(node)
	}
	node.updateHeight()
	return node
}

func (n *Node[Key, Value]) updateHeight() {
	switch {
	case n.IsLeaf():
		n.height = 1
	case n.Left == nil:
		n.height = n.Right.height + 1
	case n.Right == nil:
		n.height = n.Left.height + 1
	default:
		n.height = max(n.Left.height, n.Right.height) + 1
	}
}

func (bt *BinTree[Key, Value]) Get(key Key) (Value, bool) {
	return bt.get(bt.head, key)
}

func (bt *BinTree[Key, Value]) get(node *Node[Key, Value], key Key) (Value, bool) {
	if node == nil {
		var zero Value
		return zero, false
	}
	cmp := bt.cmp(node.Key, key)
	switch {
	case cmp > 0:
		return bt.get(node.Left, key)
	case cmp < 0:
		return bt.get(node.Right, key)
	default:
		return node.Value, true
	}
}

type TraversalAction[Key, Value any] func(Key, Value)
type TraversalActionNode[Key, Value any] func(*Node[Key, Value])

func (bt *BinTree[Key, Value]) PreOrderTraversal(
	ta TraversalAction[Key, Value],
) {
	bt.preOrderTraversal(bt.head, ta)
}

func (bt *BinTree[Key, Value]) preOrderTraversal(
	node *Node[Key, Value], ta TraversalAction[Key, Value],
) {
	if node == nil {
		return
	}
	ta(node.Key, node.Value)
	bt.preOrderTraversal(node.Left, ta)
	bt.preOrderTraversal(node.Right, ta)
}

func (bt *BinTree[Key, Value]) PreOrderTraversalNode(
	ta TraversalActionNode[Key, Value],
) {
	bt.preOrderTraversalNode(bt.head, ta)
}

func (bt *BinTree[Key, Value]) preOrderTraversalNode(
	node *Node[Key, Value], ta TraversalActionNode[Key, Value],
) {
	if node == nil {
		return
	}
	ta(node)
	bt.preOrderTraversalNode(node.Left, ta)
	bt.preOrderTraversalNode(node.Right, ta)
}

func (bt *BinTree[Key, Value]) Delete(key Key) {
	bt.head = bt.delete(bt.head, key)
}

func (bt *BinTree[Key, Value]) delete(node *Node[Key, Value], key Key) *Node[Key, Value] {
	if node == nil {
		return nil
	}
	cmp := bt.cmp(node.Key, key)
	switch {
	case cmp > 0:
		node.Left = bt.delete(node.Left, key)
	case cmp < 0:
		node.Right = bt.delete(node.Right, key)
	default:
		node = bt.deleteAt(node)
	}
	return node
}

func (bt *BinTree[Key, Value]) deleteAt(node *Node[Key, Value]) *Node[Key, Value] {
	switch {
	case node.Left == nil:
		bt.count--
		return node.Right
	case node.Right == nil:
		bt.count--
		return node.Left
	default:
		nodeMin := node.Right.minNode()
		node.Value = nodeMin.Value
		node.Key = nodeMin.Key
		node.Right = bt.delete(node.Right, nodeMin.Key)
		return node
	}
}

func (bt *BinTree[Key, Value]) balance(node *Node[Key, Value]) *Node[Key, Value] {
	if node.Left == nil || node.Right.height > node.Left.height {
		if node.Right.getBF() < 0 {
			node.Right = node.Right.rightRotation()
		}
		return node.leftRotation()
	}
	if node.Right == nil || node.Left.height > node.Right.height {
		if node.Left.getBF() > 0 {
			node.Left = node.Left.leftRotation()
		}
		return node.rightRotation()
	}
	return node
}

func (n *Node[Key, Value]) getBF() int {
	leftHeight := 0
	if n.Left != nil {
		leftHeight = n.Left.height
	}
	rightHeight := 0
	if n.Right != nil {
		rightHeight = n.Right.height
	}
	return rightHeight - leftHeight
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
