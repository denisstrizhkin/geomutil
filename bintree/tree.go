package bintree

import (
	"fmt"
	"strings"
)

type Node[Key, Value any] struct {
	Key   Key
	Value Value
	Left  *Node[Key, Value]
	Right *Node[Key, Value]
	BF    int
}

func NewNode[Key, Value any](key Key, val Value) *Node[Key, Value] {
	return &Node[Key, Value]{Key: key, Value: val, BF: 0}
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
	return fmt.Sprintf("(%v %v)", n.Key, n.Value)
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
	sb.WriteString("[ ")
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
	ly := y.Left
	y.Left = n
	n.Right = ly
	return y
}

func (n *Node[Key, Value]) rightRotation() *Node[Key, Value] {
	y := n.Left
	ry := y.Right
	y.Right = n
	n.Left = ry
	return y
}

func (bt *BinTree[Key, Value]) Put(key Key, val Value) {
	bt.head = bt.put(bt.head, key, val)
	if !bt.IsBalanced() {
		bt.head = bt.balance(bt.head, key)
	}
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
	return node
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

func (bt *BinTree[Key, Value]) balance(node *Node[Key, Value], key Key) *Node[Key, Value] {
	bt.updateBF(bt.head)
	cmp := bt.cmp(key, node.Left.Key)
	switch {
	case node.BF > 1 && cmp < 0:
		return node.rightRotation()
	case node.BF < -1 && cmp > 0:
		return node.leftRotation()
	case node.BF > 1 && cmp > 0:
		node.Left = node.leftRotation()
		return node.rightRotation()
	case node.BF < -1 && cmp < 0:
		node.Right = node.rightRotation()
		return node.leftRotation()
	}
	return node
}

func (bt *BinTree[Key, Value]) updateBF(node *Node[Key, Value]) {
	node.BF = bt.getBF(node)
	if !node.IsLeaf() {
		switch {
		case node.Left == nil:
			bt.updateBF(node.Right)
		case node.Right == nil:
			bt.updateBF(node.Left)
		default:
			bt.updateBF(node.Left)
			bt.updateBF(node.Right)
		}
	}
}

func (bt *BinTree[Key, Value]) getBF(node *Node[Key, Value]) int {
	if node == nil {
		return 0
	}
	lh := bt.getHeight(node.Left)
	rh := bt.getHeight(node.Right)
	return lh - rh
}

func (bt *BinTree[Key, Value]) getHeight(node *Node[Key, Value]) int {
	if node == nil {
		return 0
	}
	height := 1
	switch {
	case node.IsLeaf():
		return height
	case node.Left == nil:
		height += bt.getHeight(node.Right)
	case node.Right == nil:
		height += bt.getHeight(node.Left)
	default:
		height += max(bt.getHeight(node.Left), bt.getHeight(node.Right))
	}
	return height
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (bt *BinTree[Key, Value]) isBalTree(node *Node[Key, Value]) int {
	if node == nil {
		return 0
	}
	lh := bt.isBalTree(node.Left)
	if lh == -1 {
		return -1
	}
	rh := bt.isBalTree(node.Right)
	if rh == -1 {
		return -1
	}

	if absInt(lh-rh) > 1 {
		return -1
	}
	return max(lh, rh) + 1
}

func (bt *BinTree[Key, Value]) IsBalanced() bool {
	return bt.isBalTree(bt.head) > 0
}
