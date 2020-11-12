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

	if timetable.First().subject != timetable._1.subject {
		t.Fatal("First() should return 1st class")
	}
	if timetable.Second().subject != timetable._2.subject {
		t.Fatal("Second() should return 2nd class")
	}
	if timetable.Third().subject != timetable._3.subject {
		t.Fatal("Third() should return 3rd class")
	}
	if timetable.Fourth().subject != timetable._4.subject {
		t.Fatal("Fourth() should return 4th class")
	}
	if timetable.Fifth().subject != timetable._5.subject {
		t.Fatal("Fifth() should return 5th class")
	}
}
