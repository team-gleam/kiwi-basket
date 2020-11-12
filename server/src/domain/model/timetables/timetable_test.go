package timetables

import (
	"testing"
)

func TestTimetableGetters(t *testing.T) {
	timetable := NewTimetable(
		NoRoom("1", ""),
		NoRoom("2", ""),
		NoRoom("3", ""),
		NoRoom("4", ""),
		NoRoom("5", ""),
	)

	if timetable.First().subject != "1" {
		t.Fatal("First() should return 1st class")
	}
	if timetable.Second().subject != "2" {
		t.Fatal("Second() should return 2nd class")
	}
	if timetable.Third().subject != "3" {
		t.Fatal("Third() should return 3rd class")
	}
	if timetable.Fourth().subject != "4" {
		t.Fatal("Fourth() should return 4th class")
	}
	if timetable.Fifth().subject != "5" {
		t.Fatal("Fifth() should return 5th class")
	}
}
