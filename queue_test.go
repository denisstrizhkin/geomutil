package geomutil

import (
	"fmt"
	"testing"
)

func strWantedGot(msg string, wanted, got any) string {
	return fmt.Sprintf("%s\nwanted: %v\ngot:    %v", msg, wanted, got)
}

func TestPut(t *testing.T) {
	eq := NewEventQueue[int]()
	eq.Put(1)
	eq.Put(2)
	eq.Put(3)
	eq.Put(4)
	eq.Put(5)
	wanted := "[1 2 3 4 5]"
	if eq.String() != wanted {
		t.Error(strWantedGot("", wanted, eq))
	}
}

func TestPick(t *testing.T) {
	eq := NewEventQueue[int]()
	eq.Put(1)
	eq.Put(2)
	eq.Put(3)
	eq.Put(4)
	eq.Put(5)
	wanted := "5"
	if fmt.Sprint(eq.Pick()) != wanted {
		t.Error(strWantedGot("", wanted, fmt.Sprint(eq.Pick())))
	}
}

func TestPop(t *testing.T) {
	eq := NewEventQueue[int]()
	eq.Put(1)
	eq.Put(2)
	eq.Put(3)
	eq.Put(4)
	eq.Put(5)
	wanted := "5"
	if fmt.Sprint(eq.Pop()) != wanted {
		t.Error(strWantedGot("", wanted, fmt.Sprint(eq.Pop())))
	}
	wanted = "[1 2 3 4]"
	if eq.String() != wanted {
		t.Error(strWantedGot("", wanted, eq.String()))
	}
}
