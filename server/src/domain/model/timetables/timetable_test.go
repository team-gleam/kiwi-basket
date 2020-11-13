package timetables

import (
	"testing"
)

var (
	_1 = NoRoom("1", "")
	_2 = NoRoom("2", "")
	_3 = NoRoom("3", "")
	_4 = NoRoom("4", "")
	_5 = NoRoom("5", "")
)

func TestNewTimetable(t *testing.T) {
	timetable := NewTimetable(
		_1,
		_2,
		_3,
		_4,
		_5,
	)

	tests := []struct {
		got      string
		expected string
	}{
		{timetable._1.subject, _1.subject},
		{timetable._2.subject, _2.subject},
		{timetable._3.subject, _3.subject},
		{timetable._4.subject, _4.subject},
		{timetable._5.subject, _5.subject},
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

func TestTimetableGetters(t *testing.T) {
	timetable := Timetable{
		_1,
		_2,
		_3,
		_4,
		_5,
	}

	tests := []struct {
		got      string
		expected string
	}{
		{timetable.First().subject, _1.subject},
		{timetable.Second().subject, _2.subject},
		{timetable.Third().subject, _3.subject},
		{timetable.Fourth().subject, _4.subject},
		{timetable.Fifth().subject, _5.subject},
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
