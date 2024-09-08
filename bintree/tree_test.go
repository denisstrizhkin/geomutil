package bintree

import (
	// "fmt"
	"testing"
)

func cmpInt(a, b int) int {
	return a - b
}

func TestPut(t *testing.T) {
	//     _3
	//    /  \
	//   1   4
	//  / \   \
	// 0  2   5
	var bt *BinTree[int, string]
	cases := []struct {
		Action func()
		State  string
		Size   int
	}{
		{
			func() { bt = NewBinTree[int, string](cmpInt) },
			"[ ]", 0,
		}, {
			func() { bt.Put(3, "Mion") },
			"[ (3 Mion) ]", 1,
		}, {
			func() { bt.Put(1, "Shion") },
			"[ (3 Mion) (1 Shion) ]", 2,
		}, {
			func() { bt.Put(4, "Misaki") },
			"[ (3 Mion) (1 Shion) (4 Misaki) ]", 3,
		}, {
			func() { bt.Put(0, "Rena") },
			"[ (3 Mion) (1 Shion) (0 Rena) (4 Misaki) ]", 4,
		}, {
			func() { bt.Put(5, "Sakura") },
			"[ (3 Mion) (1 Shion) (0 Rena) (4 Misaki) (5 Sakura) ]", 5,
		}, {
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

func TestGet(t *testing.T) {
	//     _3
	//    /  \
	//   1   4
	//  /
	// 0
	bt := NewBinTree[int, string](cmpInt)
	bt.Put(3, "Mion")
	bt.Put(1, "Shion")
	bt.Put(4, "Misaki")
	bt.Put(0, "Rena")
	cases := []struct {
		Action func() (string, bool)
		Want   string
		Ok     bool
	}{
		{func() (string, bool) { return bt.Get(3) }, "Mion", true},
		{func() (string, bool) { return bt.Get(1) }, "Shion", true},
		{func() (string, bool) { return bt.Get(4) }, "Misaki", true},
		{func() (string, bool) { return bt.Get(0) }, "Rena", true},
		{func() (string, bool) { return bt.Get(10) }, "", false},
	}
	stateWant := "[ (3 Mion) (1 Shion) (0 Rena) (4 Misaki) ]"
	sizeWant := 4
	for _, testCase := range cases {
		got, ok := testCase.Action()
		if testCase.Want != got || testCase.Ok != ok {
			t.Errorf("wanted: %s | got: %s", testCase.Want, got)
		}
		state := bt.String()
		size := bt.Size()
		if stateWant != state || sizeWant != size {
			t.Errorf(
				"\nwanted | size: %d, state: %s\n   got | size: %d, state: %s",
				sizeWant, stateWant, size, state,
			)
		}
	}
}

func TestNodeString(t *testing.T) {
	n := NewNode(1, "Rika")
	if n.String() != "(1 Rika)" {
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
	wanted := "[ (3 3) (2 2) (1 1) (4 4) ]"
	if bt.String() != wanted {
		t.Errorf("wanted %s, got %v", wanted, bt)
	}
}

func TestNodeDelete(t *testing.T) {
	bt := NewBinTree[int, int](cmpInt)
	bt.Put(3, 3)
	bt.Put(2, 2)
	bt.Put(4, 4)
	bt.Put(1, 1)
	bt.Delete(1)
	wanted := "[ (3 3) (2 2) (4 4) ]"
	if bt.String() != wanted {
		t.Errorf("wanted %s, got %v", wanted, bt)
	}

	bt1 := NewBinTree[int, int](cmpInt)
	bt1.Put(3, 3)
	bt1.Put(2, 2)
	bt1.Put(4, 4)
	bt1.Put(1, 1)
	bt1.Delete(3)
	wanted = "[ (4 4) (2 2) (1 1) ]"
	if bt1.String() != wanted {
		t.Errorf("wanted %s, got %v", wanted, bt1)
	}
}
