package timetables

import (
	"testing"

	timetablesModel "github.com/team-gleam/kiwi-basket/domain/model/timetables"
)

type TimetablesResponseToTimetables struct {
	Name     string
	Input    TimetablesResponse
	Expected timetablesModel.Timetables
}

var t1 = TimetablesResponseToTimetables{
	Name: "no null",
	Input: TimetablesResponse{
		Timetables: Timetables{
			Timetable{
				One:   &Class{"A", "1"},
				Two:   &Class{"B", "2"},
				Three: &Class{"C", "3"},
				Four:  &Class{"D", "4"},
				Five:  &Class{"E", "5"},
			},
			Timetable{
				One:   &Class{"F", "6"},
				Two:   &Class{"G", "7"},
				Three: &Class{"H", "8"},
				Four:  &Class{"I", "9"},
				Five:  &Class{"J", "10"},
			},
			Timetable{
				One:   &Class{"K", "11"},
				Two:   &Class{"L", "12"},
				Three: &Class{"M", "13"},
				Four:  &Class{"N", "14"},
				Five:  &Class{"O", "15"},
			},
			Timetable{
				One:   &Class{"P", "16"},
				Two:   &Class{"Q", "17"},
				Three: &Class{"R", "18"},
				Four:  &Class{"S", "19"},
				Five:  &Class{"T", "20"},
			},
			Timetable{
				One:   &Class{"U", "21"},
				Two:   &Class{"V", "22"},
				Three: &Class{"W", "23"},
				Four:  &Class{"X", "24"},
				Five:  &Class{"Y", "25"},
			},
		},
	},
	Expected: timetablesModel.NewTimetables(
		timetablesModel.NewTimetable(
			timetablesModel.NewClass("A", "1"),
			timetablesModel.NewClass("B", "2"),
			timetablesModel.NewClass("C", "3"),
			timetablesModel.NewClass("D", "4"),
			timetablesModel.NewClass("E", "5"),
		),
		timetablesModel.NewTimetable(
			timetablesModel.NewClass("F", "6"),
			timetablesModel.NewClass("G", "7"),
			timetablesModel.NewClass("H", "8"),
			timetablesModel.NewClass("I", "9"),
			timetablesModel.NewClass("J", "10"),
		),
		timetablesModel.NewTimetable(
			timetablesModel.NewClass("K", "11"),
			timetablesModel.NewClass("L", "12"),
			timetablesModel.NewClass("M", "13"),
			timetablesModel.NewClass("N", "14"),
			timetablesModel.NewClass("O", "15"),
		),
		timetablesModel.NewTimetable(
			timetablesModel.NewClass("P", "16"),
			timetablesModel.NewClass("Q", "17"),
			timetablesModel.NewClass("R", "18"),
			timetablesModel.NewClass("S", "19"),
			timetablesModel.NewClass("T", "20"),
		),
		timetablesModel.NewTimetable(
			timetablesModel.NewClass("U", "21"),
			timetablesModel.NewClass("V", "22"),
			timetablesModel.NewClass("W", "23"),
			timetablesModel.NewClass("X", "24"),
			timetablesModel.NewClass("Y", "25"),
		),
	),
}

var t2 = TimetablesResponseToTimetables{
	Name: "has null",
	Input: TimetablesResponse{
		Timetables: Timetables{
			Timetable{
				One:   nil,
				Two:   &Class{"B", "2"},
				Three: &Class{"C", "3"},
				Four:  &Class{"D", "4"},
				Five:  &Class{"E", "5"},
			},
			Timetable{
				One:   &Class{"F", "6"},
				Two:   nil,
				Three: &Class{"H", "8"},
				Four:  &Class{"I", "9"},
				Five:  &Class{"J", "10"},
			},
			Timetable{
				One:   &Class{"K", "11"},
				Two:   &Class{"L", "12"},
				Three: nil,
				Four:  &Class{"N", "14"},
				Five:  &Class{"O", "15"},
			},
			Timetable{
				One:   &Class{"P", "16"},
				Two:   &Class{"Q", "17"},
				Three: &Class{"R", "18"},
				Four:  nil,
				Five:  &Class{"T", "20"},
			},
			Timetable{
				One:   &Class{"U", "21"},
				Two:   &Class{"V", "22"},
				Three: &Class{"W", "23"},
				Four:  &Class{"X", "24"},
				Five:  nil,
			},
		},
	},
	Expected: timetablesModel.NewTimetables(
		timetablesModel.NewTimetable(
			timetablesModel.NoClass(),
			timetablesModel.NewClass("B", "2"),
			timetablesModel.NewClass("C", "3"),
			timetablesModel.NewClass("D", "4"),
			timetablesModel.NewClass("E", "5"),
		),
		timetablesModel.NewTimetable(
			timetablesModel.NewClass("F", "6"),
			timetablesModel.NoClass(),
			timetablesModel.NewClass("H", "8"),
			timetablesModel.NewClass("I", "9"),
			timetablesModel.NewClass("J", "10"),
		),
		timetablesModel.NewTimetable(
			timetablesModel.NewClass("K", "11"),
			timetablesModel.NewClass("L", "12"),
			timetablesModel.NoClass(),
			timetablesModel.NewClass("N", "14"),
			timetablesModel.NewClass("O", "15"),
		),
		timetablesModel.NewTimetable(
			timetablesModel.NewClass("P", "16"),
			timetablesModel.NewClass("Q", "17"),
			timetablesModel.NewClass("R", "18"),
			timetablesModel.NoClass(),
			timetablesModel.NewClass("T", "20"),
		),
		timetablesModel.NewTimetable(
			timetablesModel.NewClass("U", "21"),
			timetablesModel.NewClass("V", "22"),
			timetablesModel.NewClass("W", "23"),
			timetablesModel.NewClass("X", "24"),
			timetablesModel.NoClass(),
		),
	),
}

func TestTimetablesResponseToTimetables(t *testing.T) {
	tcs := []TimetablesResponseToTimetables{t1, t2}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			tt := tc.Input.toTimetables()
			if tt != tc.Expected {
				t.Errorf("Failed# expected: %v; got: %v\n", tc.Expected, tt)
			}
		})
	}
}
