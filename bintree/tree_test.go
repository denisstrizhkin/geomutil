package bintree

import (
	"testing"
)

func cmpInt(a, b int) int {
	return a - b
}

type TestCase struct {
	Action func()
	State  string
	Size   int
}

func TestInsert(t *testing.T) {
	//     _3
	//    /  \
	//   1   4
	//  / \   \
	// 0  2   5

	var bt *BinTree[int, string]
	cases := []TestCase{
		TestCase{
			func() { bt = NewBinTree[int, string](cmpInt) },
			"[ ]", 0,
		}, TestCase{
			func() { bt.Put(3, "Mion") },
			"[ (3 Mion) ]", 1,
		}, TestCase{
			func() { bt.Put(1, "Shion") },
			"[ (3 Mion) (1 Shion) ]", 2,
		}, TestCase{
			func() { bt.Put(4, "Misaki") },
			"[ (3 Mion) (1 Shion) (4 Misaki) ]", 3,
		}, TestCase{
			func() { bt.Put(0, "Rena") },
			"[ (3 Mion) (1 Shion) (0 Rena) (4 Misaki) ]", 4,
		}, TestCase{
			func() { bt.Put(5, "Sakura") },
			"[ (3 Mion) (1 Shion) (0 Rena) (4 Misaki) (5 Sakura) ]", 5,
		}, TestCase{
			func() { bt.Put(2, "Rosa") },
			"[ (3 Mion) (1 Shion) (0 Rena) (2 Rosa) (4 Misaki) (5 Sakura) ]", 6,
		},
	}

	for _, testCase := range cases {
		testCase.Action()
		state := bt.String()
		size := bt.Size()
		if testCase.State != state || testCase.Size != size {
			t.Errorf(
				"\nwanted | size: %d, state: %s\n   got | size: %d, state: %s",
				testCase.Size, testCase.State, size, state,
			)
		}
	}
}

// func TestNodeString(t *testing.T) {
// 	n := NewNode(1, "Rika")
// 	if n.String() != "(1 Rika) " {
// 		t.Errorf("wanted %s, got %s", "(1 Rika)", n.String())
// 	}
// }

// func TestBinTreeString(t *testing.T) {
// 	bt := NewBinTree[int, int](cmpInt)
// 	bt.Put(3, 3)
// 	bt.Put(2, 2)
// 	bt.Put(4, 4)
// 	bt.Put(1, 1)
// 	if !bt.IsBalanced() {
// 		t.Errorf("Your tree is not balanced!")
// 	}
// 	wanted := "[ (3 3) (2 2) (1 1) (4 4) ]"
// 	if bt.String() != wanted {
// 		t.Errorf("wanted %s, got %v", wanted, bt)
// 	}
// }
