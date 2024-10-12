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
	eq.Enqueue(1)
	eq.Enqueue(2)
	eq.Enqueue(3)
	eq.Enqueue(4)
	eq.Enqueue(5)
	wanted := "[1 2 3 4 5]"
	if eq.String() != wanted {
		t.Error(strWantedGot("", wanted, eq))
	}
}

func TestPick(t *testing.T) {
	eq := NewEventQueue[int]()
	eq.Enqueue(1)
	eq.Enqueue(2)
	eq.Enqueue(3)
	eq.Enqueue(4)
	eq.Enqueue(5)
	wanted := "5"
	if fmt.Sprint(eq.Peek()) != wanted {
		t.Error(strWantedGot("", wanted, fmt.Sprint(eq.Peek())))
	}
}

func TestPop(t *testing.T) {
	eq := NewEventQueue[int]()
	eq.Enqueue(1)
	eq.Enqueue(2)
	eq.Enqueue(3)
	eq.Enqueue(4)
	eq.Enqueue(5)
	wanted := "5"
	if fmt.Sprint(eq.Dequeue()) != wanted {
		t.Error(strWantedGot("", wanted, fmt.Sprint(eq.Dequeue())))
	}
	wanted = "[1 2 3 4]"
	if eq.String() != wanted {
		t.Error(strWantedGot("", wanted, eq.String()))
	}
}
