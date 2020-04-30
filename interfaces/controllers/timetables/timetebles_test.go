package timetables

import (
	"testing"

	timetablesModel "github.com/team-gleam/kiwi-basket/domain/model/timetables"
)

var noNullTimetablesResponse = TimetablesResponse{
	Timetables: TimetablesJSON{
		TimetableJSON{
			One:   &ClassJSON{"A", "1"},
			Two:   &ClassJSON{"B", "2"},
			Three: &ClassJSON{"C", "3"},
			Four:  &ClassJSON{"D", "4"},
			Five:  &ClassJSON{"E", "5"},
		},
		TimetableJSON{
			One:   &ClassJSON{"F", "6"},
			Two:   &ClassJSON{"G", "7"},
			Three: &ClassJSON{"H", "8"},
			Four:  &ClassJSON{"I", "9"},
			Five:  &ClassJSON{"J", "10"},
		},
		TimetableJSON{
			One:   &ClassJSON{"K", "11"},
			Two:   &ClassJSON{"L", "12"},
			Three: &ClassJSON{"M", "13"},
			Four:  &ClassJSON{"N", "14"},
			Five:  &ClassJSON{"O", "15"},
		},
		TimetableJSON{
			One:   &ClassJSON{"P", "16"},
			Two:   &ClassJSON{"Q", "17"},
			Three: &ClassJSON{"R", "18"},
			Four:  &ClassJSON{"S", "19"},
			Five:  &ClassJSON{"T", "20"},
		},
		TimetableJSON{
			One:   &ClassJSON{"U", "21"},
			Two:   &ClassJSON{"V", "22"},
			Three: &ClassJSON{"W", "23"},
			Four:  &ClassJSON{"X", "24"},
			Five:  &ClassJSON{"Y", "25"},
		},
	},
}

var noNullTimetables = timetablesModel.NewTimetables(
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
)

var hasNullTimetablesResponse = TimetablesResponse{
	Timetables: TimetablesJSON{
		TimetableJSON{
			One:   nil,
			Two:   &ClassJSON{"B", "2"},
			Three: &ClassJSON{"C", "3"},
			Four:  &ClassJSON{"D", "4"},
			Five:  &ClassJSON{"E", "5"},
		},
		TimetableJSON{
			One:   &ClassJSON{"F", "6"},
			Two:   nil,
			Three: &ClassJSON{"H", "8"},
			Four:  &ClassJSON{"I", "9"},
			Five:  &ClassJSON{"J", "10"},
		},
		TimetableJSON{
			One:   &ClassJSON{"K", "11"},
			Two:   &ClassJSON{"L", "12"},
			Three: nil,
			Four:  &ClassJSON{"N", "14"},
			Five:  &ClassJSON{"O", "15"},
		},
		TimetableJSON{
			One:   &ClassJSON{"P", "16"},
			Two:   &ClassJSON{"Q", "17"},
			Three: &ClassJSON{"R", "18"},
			Four:  nil,
			Five:  &ClassJSON{"T", "20"},
		},
		TimetableJSON{
			One:   &ClassJSON{"U", "21"},
			Two:   &ClassJSON{"V", "22"},
			Three: &ClassJSON{"W", "23"},
			Four:  &ClassJSON{"X", "24"},
			Five:  nil,
		},
	},
}

var hasNullTimetables = timetablesModel.NewTimetables(
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
)

type TimetablesResponseToTimetables struct {
	Name     string
	Input    TimetablesResponse
	Expected timetablesModel.Timetables
}

var tr2t1 = TimetablesResponseToTimetables{
	Name:     "no null",
	Input:    noNullTimetablesResponse,
	Expected: noNullTimetables,
}

var tr2t2 = TimetablesResponseToTimetables{
	Name:     "has null",
	Input:    hasNullTimetablesResponse,
	Expected: hasNullTimetables,
}

func TestTimetablesResponseToTimetables(t *testing.T) {
	tcs := []TimetablesResponseToTimetables{tr2t1, tr2t2}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			tt := tc.Input.toTimetables()
			if tt != tc.Expected {
				t.Errorf("Failed# expected: %v; got: %v\n", tc.Expected, tt)
			}
		})
	}
}

type TimetablesToTimetablesResponse struct {
	Name     string
	Input    timetablesModel.Timetables
	Expected TimetablesResponse
}

var t2tr1 = TimetablesToTimetablesResponse{
	Name:     "no null",
	Input:    noNullTimetables,
	Expected: noNullTimetablesResponse,
}

var t2tr2 = TimetablesToTimetablesResponse{
	Name:     "has null",
	Input:    hasNullTimetables,
	Expected: hasNullTimetablesResponse,
}

func TestTimetablesToTimetablesResponse(t *testing.T) {
	tcs := []TimetablesToTimetablesResponse{t2tr1, t2tr2}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			tt := toTimetablesResponse(tc.Input)
			if tt.toTimetables() != tc.Expected.toTimetables() {
				t.Errorf("Failed# expected: %v; got: %v\n", tc.Expected, tt)
			}
		})
	}
}
