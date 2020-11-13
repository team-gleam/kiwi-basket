package timetables

import (
	"testing"
)

var (
	mon = Timetable{
		NoRoom("1", ""),
		NoClass(),
		NoClass(),
		NoClass(),
		NoClass(),
	}
	tue = Timetable{
		NoRoom("2", ""),
		NoClass(),
		NoClass(),
		NoClass(),
		NoClass(),
	}
	wed = Timetable{
		NoRoom("3", ""),
		NoClass(),
		NoClass(),
		NoClass(),
		NoClass(),
	}
	thu = Timetable{
		NoRoom("4", ""),
		NoClass(),
		NoClass(),
		NoClass(),
		NoClass(),
	}
	fri = Timetable{
		NoRoom("5", ""),
		NoClass(),
		NoClass(),
		NoClass(),
		NoClass(),
	}
)

func TestNewTimetables(t *testing.T) {
	timetables := NewTimetables(
		mon,
		tue,
		wed,
		thu,
		fri,
	)

	tests := []struct {
		expected string
		got      string
	}{
		{mon._1.subject, timetables.mon._1.subject},
		{tue._1.subject, timetables.tue._1.subject},
		{wed._1.subject, timetables.wed._1.subject},
		{thu._1.subject, timetables.thu._1.subject},
		{fri._1.subject, timetables.fri._1.subject},
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

func TestTimetablesGetters(t *testing.T) {
	timetables := Timetables{
		mon,
		tue,
		wed,
		thu,
		fri,
	}

	tests := []struct {
		got      string
		expected string
	}{
		{timetables.Mon()._1.subject, mon._1.subject},
		{timetables.Tue()._1.subject, tue._1.subject},
		{timetables.Wed()._1.subject, wed._1.subject},
		{timetables.Thu()._1.subject, thu._1.subject},
		{timetables.Fri()._1.subject, fri._1.subject},
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
