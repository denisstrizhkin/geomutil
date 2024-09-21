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

func (n *Node[Key, Value]) String() string {
	if n == nil {
		return fmt.Sprint(nil)
	}
	return fmt.Sprintf("(%v| %v: %v)", n.height, n.Key, n.Value)
}

func (n *Node[Key, Value]) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

type Comparator[Key any] func(a, b Key) int

type BinTree[Key, Value any] struct {
	head *Node[Key, Value]
	cmp  Comparator[Key]
}

func NewBinTree[Key, Value any](comp Comparator[Key]) *BinTree[Key, Value] {
	return &BinTree[Key, Value]{cmp: comp}
}

func (bt *BinTree[Key, Value]) String() string {
	if bt == nil {
		return fmt.Sprint(nil)
	}
	var sb strings.Builder
	sb.WriteString("[ ")
	bt.PreOrderTraversalNode(func(node *Node[Key, Value]) {
		sb.WriteString(node.String())
		sb.WriteRune(' ')
	})
	sb.WriteRune(']')
	return sb.String()
}

func (n *Node[Key, Value]) minNode() *Node[Key, Value] {
	if n == nil {
		return nil
	}
	if n.Left == nil {
		return n
	}
	return n.Left.minNode()
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
	bt.head = bt.head.put(key, val, bt.cmp)
}

func (n *Node[Key, Value]) put(
	key Key, val Value, cmp Comparator[Key],
) *Node[Key, Value] {
	if n == nil {
		return NewNode(key, val)
	}
	c := cmp(n.Key, key)
	switch {
	case c > 0:
		n.Left = n.Left.put(key, val, cmp)
	case c < 0:
		n.Right = n.Right.put(key, val, cmp)
	default:
		n.Value = val
	}
	n.updateHeight()
	return n.balance()
}

func (n *Node[Key, Value]) updateHeight() {
	if n != nil {
		n.height = max(n.Left.getHeight(), n.Right.getHeight()) + 1
	}
}

func (bt *BinTree[Key, Value]) Get(key Key) (Value, bool) {
	return bt.head.get(key, bt.cmp)
}

func (n *Node[Key, Value]) get(key Key, cmp Comparator[Key]) (Value, bool) {
	if n == nil {
		var zero Value
		return zero, false
	}
	c := cmp(n.Key, key)
	switch {
	case c > 0:
		return n.Left.get(key, cmp)
	case c < 0:
		return n.Right.get(key, cmp)
	default:
		return n.Value, true
	}
}

type TraversalAction[Key, Value any] func(Key, Value)
type TraversalActionNode[Key, Value any] func(*Node[Key, Value])

func (bt *BinTree[Key, Value]) PreOrderTraversal(
	ta TraversalAction[Key, Value],
) {
	bt.head.preOrderTraversal(ta)
}

func (n *Node[Key, Value]) preOrderTraversal(
	ta TraversalAction[Key, Value],
) {
	if n == nil {
		return
	}
	ta(n.Key, n.Value)
	n.Left.preOrderTraversal(ta)
	n.Left.preOrderTraversal(ta)
}

func (bt *BinTree[Key, Value]) PreOrderTraversalNode(
	ta TraversalActionNode[Key, Value],
) {
	bt.head.preOrderTraversalNode(ta)
}

func (n *Node[Key, Value]) preOrderTraversalNode(
	ta TraversalActionNode[Key, Value],
) {
	if n == nil {
		return
	}
	ta(n)
	n.Left.preOrderTraversalNode(ta)
	n.Right.preOrderTraversalNode(ta)
}

func (bt *BinTree[Key, Value]) Delete(key Key) {
	bt.head = bt.head.delete(key, bt.cmp)
}

func (n *Node[Key, Value]) delete(key Key, cmp Comparator[Key]) *Node[Key, Value] {
	if n == nil {
		return nil
	}
	c := cmp(n.Key, key)
	switch {
	case c > 0:
		n.Left = n.Left.delete(key, cmp)
	case c < 0:
		n.Right = n.Right.delete(key, cmp)
	default:
		switch {
		case n.Left == nil:
			n = n.Right
		case n.Right == nil:
			n = n.Left
		default:
			nodeMin := n.Right.minNode()
			n.Value = nodeMin.Value
			n.Key = nodeMin.Key
			n.Right = n.Right.delete(nodeMin.Key, cmp)
		}
	}
	n.updateHeight()
	return n.balance()
}

func (n *Node[Key, Value]) balance() *Node[Key, Value] {
	bf := n.getBF()
	if bf > 1 {
		if n.Right.getBF() < 0 {
			n.Right = n.Right.rightRotation()
		}
		return n.leftRotation()
	}
	if bf < -1 {
		if n.Left.getBF() > 0 {
			n.Left = n.Left.leftRotation()
		}
		return n.rightRotation()
	}
	return n
}

func (n *Node[Key, Value]) getBF() int {
	if n == nil {
		return 0
	}
	return n.Right.getHeight() - n.Left.getHeight()
}
