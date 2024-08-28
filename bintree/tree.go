package bintree

type BinTree[Key, Value any] struct {
	head  *Node[Key, Value]
	cmp   Comparator[Key]
	count int
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

type Comparator[Key any] func(a, b Key) int

func NewBinTree[Key, Value any](comp Comparator[Key]) *BinTree[Key, Value] {
	return &BinTree[Key, Value]{cmp: comp}
}

func (bt *BinTree[Key, Value]) Size() int {
	return bt.count
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
	case cmp < 0:
		return bt.put(node.Left, key, val)
	case cmp > 0:
		return bt.put(node.Right, key, val)
	default:
		node.Value = val
		return node
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
	case cmp < 0:
		return bt.get(node.Left, key)
	case cmp > 0:
		return bt.get(node.Right, key)
	default:
		return node.Value, true
	}
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
	case cmp < 0:
		return bt.delete(node.Left, key)
	case cmp > 0:
		return bt.delete(node.Right, key)
	default:
		bt.count--
		return bt.deleteAt(node)
	}
}

func (bt *BinTree[Key, Value]) deleteAt(node *Node[Key, Value]) *Node[Key, Value] {
	switch {
	case node.Left == nil:
		return node.Right
	case node.Right == nil:
		return node.Left
	default:
		nodeMin := node.Right.minNode()
		node.Value = nodeMin.Value
		bt.deleteAt(nodeMin)
		return node
	}
}
