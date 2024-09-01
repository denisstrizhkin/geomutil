package bintree

import (
	"reflect"
	"testing"
)

func cmpInt(a, b int) int {
	return a - b
}

var ExTree BinTree[int, string] = BinTree[int, string]{
	count: 4,
	cmp:   cmpInt,
	head: &Node[int, string]{
		Key:   2,
		Value: "Mion",
		Left: &Node[int, string]{
			Key:   1,
			Value: "Shion",
		},
		Right: &Node[int, string]{
			Key:   3,
			Value: "Rika",
			Right: &Node[int, string]{
				Key:   4,
				Value: "Rena",
			},
		},
	},
}

func TestInsert(t *testing.T) {
	//   2
	//  / \
	// 1  3
	//   / \
	//     4
	bt := NewBinTree[int, string](cmpInt)
	if bt.Size() != 0 {
		t.Errorf("wanted %d, got %d", 0, bt.Size())
	}
	bt.Put(2, "Mion")
	if bt.Size() != 1 {
		t.Errorf("wanted %d, got %d", 0, bt.Size())
	}
	bt.Put(1, "Shion")
	if bt.Size() != 2 {
		t.Errorf("wanted %d, got %d", 0, bt.Size())
	}
	bt.Put(3, "Rika")
	if bt.Size() != 3 {
		t.Errorf("wanted %d, got %d", 0, bt.Size())
	}
	bt.Put(4, "Rena")
	if bt.Size() != 4 {
		t.Errorf("wanted %d, got %d", 0, bt.Size())
	}
	bt.Put(5, "Den4ik")
	if bt.Size() != 5 {
		t.Errorf("wanted %d, got %d", 0, bt.Size())
	}
	if reflect.DeepEqual(bt, ExTree) {
		t.Errorf("didn't match the example one")
	}
	if !bt.IsBalanced() {
		t.Errorf("Your tree is not balanced!")
	}
}

func TestNodeString(t *testing.T) {
	n := NewNode(1, "Rika")
	if n.String() != "(1 Rika) " {
		t.Errorf("wanted %s, got %s", "(1 Rika)", n.String())
	}
}

func TestBinTreeString(t *testing.T) {
	bt := NewBinTree[int, int](cmpInt)
	bt.Put(3, 3)
	bt.Put(2, 2)
	bt.Put(4, 4)
	bt.Put(1, 1)
	if !bt.IsBalanced() {
		t.Errorf("Your tree is not balanced!")
	}
	if bt.String() != "[ (3 3) (2 2) (1 1) (4 4) ]" {
		t.Errorf("wanted %s, got %s", "[ (3 3) (2 2) (1 1) (4 4) ]", bt.String())
	}
}
