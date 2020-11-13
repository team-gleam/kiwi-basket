package timetables

import (
	"testing"
)

func TestNewClass(t *testing.T) {
	class := NewClass("", "", "")
	if class.noClass || class.noRoom {
		t.Error("no class and no room flag should be false")
	}
}

func TestNoClass(t *testing.T) {
	class := NoClass()
	if !class.noClass {
		t.Error("no class flag should be true")
	}
}

func TestNoRoom(t *testing.T) {
	class := NoRoom("", "")
	if !class.noRoom {
		t.Error("no room flag should be true")
	}
}

func TestIsNoClass(t *testing.T) {
	class := NoClass()
	if !class.IsNoClass() {
		t.Error("should be true if noClass is given")
	}
}

func TestIsNoRoom(t *testing.T) {
	class := NoRoom("", "")
	if !class.IsNoRoom() {
		t.Error("should be true if noRoom is given")
	}
}

func TestClassGetters(t *testing.T) {
	class := NewClass("subject", "room", "memo")
	tests := []struct {
		expected string
		got      string
	}{
		{class.subject, class.Subject()},
		{class.room, class.Room()},
		{class.memo, class.Memo()},
	}

	for _, test := range tests {
		if test.expected != test.got {
			t.Fatalf(
				"expected: %v; got: %v\n",
				test.expected,
				test.got,
			)
		}
	}
}
