package bintree

import (
	"reflect"
	"testing"
)

func cmpInt(a, b int) int {
	return a - b
}

var exampleTree BinTree[int, string] = BinTree[int, string]{
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
	if reflect.DeepEqual(bt, exampleTree) {
		t.Errorf("didn't match the example one")
	}
}
