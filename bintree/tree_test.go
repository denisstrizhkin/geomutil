package bintree

import (
	"fmt"
	"testing"
)

func cmpInt(a, b int) int {
	return a - b
}

func strWantedGot(msg string, wanted, got any) string {
	return fmt.Sprintf("%s\nwanted: %v\ngot:    %v", msg, wanted, got)
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
	}{
		{
			func() { bt = NewBinTree[int, string](cmpInt) },
			"0| [ ]",
		}, {
			func() { bt.Put(3, "Mion") },
			"1| [ (1| 3: Mion) ]",
		}, {
			func() { bt.Put(1, "Shion") },
			"2| [ (2| 3: Mion) (1| 1: Shion) ]",
		}, {
			func() { bt.Put(4, "Misaki") },
			"3| [ (2| 3: Mion) (1| 1: Shion) (1| 4: Misaki) ]",
		}, {
			func() { bt.Put(0, "Rena") },
			"4| [ (3| 3: Mion) (2| 1: Shion) (1| 0: Rena) (1| 4: Misaki) ]",
		}, {
			func() { bt.Put(5, "Sakura") },
			"5| [ (3| 3: Mion) (2| 1: Shion) (1| 0: Rena) (2| 4: Misaki) (1| 5: Sakura) ]",
		}, {
			func() { bt.Put(2, "Rosa") },
			"6| [ (3| 3: Mion) (2| 1: Shion) (1| 0: Rena) (1| 2: Rosa) (2| 4: Misaki) (1| 5: Sakura) ]",
		},
	}
	for _, testCase := range cases {
		testCase.Action()
		if testCase.State != bt.String() {
			t.Error(strWantedGot("", testCase.State, bt))
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
	wanted := "4| [ (3| 3: Mion) (2| 1: Shion) (1| 0: Rena) (1| 4: Misaki) ]"
	for _, testCase := range cases {
		got, ok := testCase.Action()
		if testCase.Want != got || testCase.Ok != ok {
			t.Error(strWantedGot("", testCase.Want, got))
		}
		if wanted != bt.String() {
			t.Error(strWantedGot("", wanted, bt))
		}
	}
}

func TestNodeString(t *testing.T) {
	n := NewNode(1, "Rika")
	wanted := "(1| 1: Rika)"
	if n.String() != wanted {
		t.Error(strWantedGot("", wanted, n))
	}
}

func TestBinTreeString(t *testing.T) {
	bt := NewBinTree[int, string](cmpInt)
	bt.Put(30, "a")
	bt.Put(20, "d")
	bt.Put(40, "w")
	bt.Put(10, "m")
	if absInt(bt.head.getBF()) > 2 {
		t.Errorf("Your tree is not balanced!")
	}
	wanted := "4| [ (3| 30: a) (2| 20: d) (1| 10: m) (1| 40: w) ]"
	if bt.String() != wanted {
		t.Error(strWantedGot("", wanted, bt))
	}
}

func TestNodeDelete(t *testing.T) {
	bt := NewBinTree[int, int](cmpInt)
	bt.Put(3, 3)
	bt.Put(2, 2)
	bt.Put(4, 4)
	bt.Put(1, 1)
	bt.Delete(1)
	wanted := "3| [ (3 3) (2 2) (4 4) ]"
	if bt.String() != wanted {
		t.Error(strWantedGot("", wanted, bt))
	}
	if bt.count != 3 {
		t.Errorf("supposed number of elements %d, but recieved %d", 3, bt.count)
	}

	bt1 := NewBinTree[int, int](cmpInt)
	bt1.Put(3, 3)
	bt1.Put(2, 2)
	bt1.Put(4, 4)
	bt1.Put(1, 1)
	bt1.Delete(3)
	wanted = "3| [ (4 4) (2 2) (1 1) ]"
	if bt1.String() != wanted {
		t.Error(strWantedGot("", wanted, bt1))
	}
	if bt1.count != 3 {
		t.Errorf("supposed number of elements %d, but recieved %d", 3, bt1.count)
	}
}

func TestRightRotation(t *testing.T) {
	bt := NewBinTree[int, string](cmpInt)
	bt.Put(4, "z")
	bt.Put(2, "a")
	bt.Put(5, "u")
	bt.Put(1, "m")
	bt.Put(3, "h")
	{
		wanted := "5| [ (3| 4: z) (2| 2: a) (1| 1: m) (1| 3: h) (1| 5: u) ]"
		if bt.String() != wanted {
			t.Error(strWantedGot("", wanted, bt))
		}
	}
	bt.head = bt.head.rightRotation()
	{
		wanted := "5| [ (3| 2: a) (1| 1: m) (2| 4: z) (1| 3: h) (1| 5: u) ]"
		if bt.String() != wanted {
			t.Error(strWantedGot("Right rotation is incorrect", wanted, bt))
		}
	}
}

func TestLeftRotation(t *testing.T) {
	bt := NewBinTree[int, string](cmpInt)
	bt.Put(2, "z")
	bt.Put(1, "a")
	bt.Put(4, "m")
	bt.Put(3, "h")
	bt.Put(5, "u")
	{
		wanted := "5| [ (3| 2: z) (1| 1: a) (2| 4: m) (1| 3: h) (1| 5: u) ]"
		if bt.String() != wanted {
			t.Error(strWantedGot("", wanted, bt))
		}
	}
	bt.head = bt.head.leftRotation()
	{
		wanted := "5| [ (3| 4: m) (2| 2: z) (1| 1: a) (1| 3: h) (1| 5: u) ]"
		if bt.String() != wanted {
			t.Error(strWantedGot("Left rotation is incorrect", wanted, bt))
		}
	}
}

func TestTestBalanceLeft(t *testing.T) {
	bt := NewBinTree[int, string](cmpInt)
	bt.Put(2, "b")
	bt.Put(1, "a")
	bt.Put(5, "e")
	bt.Put(3, "c")
	bt.Put(6, "f")
	{
		wanted := "5| [ (3| 2: b) (1| 1: a) (2| 5: e) (1| 3: c) (1| 6: f) ]"
		if bt.String() != wanted {
			t.Error(strWantedGot("", wanted, bt))
		}
	}
	bt.Put(7, "g")
	{
		wanted := "6| [ (3| 5: e) (2| 2: b) (1| 1: a) (1| 3: c) (2| 6: f) (1| 7: g) ]"
		if bt.String() != wanted {
			t.Error(strWantedGot("", wanted, bt))
		}
	}
}

func TestTestBalanceRightLeft(t *testing.T) {
	bt := NewBinTree[int, string](cmpInt)
	bt.Put(2, "b")
	bt.Put(1, "a")
	bt.Put(5, "e")
	bt.Put(3, "c")
	bt.Put(6, "f")
	{
		wanted := "5| [ (3| 2: b) (1| 1: a) (2| 5: e) (1| 3: c) (1| 6: f) ]"
		if bt.String() != wanted {
			t.Error(strWantedGot("", wanted, bt))
		}
	}
	bt.Put(4, "d")
	{
		wanted := "6| [ (3| 3: c) (2| 2: b) (1| 1: a) (2| 5: e) (1| 4: d) (1| 6: f) ]"
		if bt.String() != wanted {
			t.Error(strWantedGot("", wanted, bt))
		}
	}
}
