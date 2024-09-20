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
	return n.minNode()
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
	if node.getBF()*node.getBF() > 1 {
		node = node.balance()
	}
	node.updateHeight()
	return node
}

func (n *Node[Key, Value]) updateHeight() {
	if n != nil {
		n.height = max(n.Left.getHeight(), n.Right.getHeight()) + 1
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
	bt.head = bt.head.delete(key, bt.cmp)
}

func (node *Node[Key, Value]) delete(key Key, cmp Comparator[Key]) *Node[Key, Value] {
	if node == nil {
		return nil
	}
	c := cmp(node.Key, key)
	switch {
	case c > 0:
		node.Left = node.Left.delete(key, cmp)
	case c < 0:
		node.Right = node.Right.delete(key, cmp)
	default:
		switch {
		case node.Left == nil:
			node = node.Right
		case node.Right == nil:
			node = node.Left
		default:
			nodeMin := node.Right.minNode()
			node.Value = nodeMin.Value
			node.Key = nodeMin.Key
			node.Right = node.Right.delete(nodeMin.Key, cmp)
		}
	}
	if node.getBF()*node.getBF() > 1 {
		node = node.balance()
	}
	node.updateHeight()
	return node
}

func (n *Node[Key, Value]) balance() *Node[Key, Value] {
	if n.getBF() > 0 {
		if n.Right.getBF() < 0 {
			n.Right = n.Right.rightRotation()
		}
		return n.leftRotation()
	}
	if n.getBF() < 0 {
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
