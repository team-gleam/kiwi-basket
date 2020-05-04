package timetables

import (
	"fmt"
	"strings"
	"testing"

	timetablesModel "github.com/team-gleam/kiwi-basket/domain/model/timetables"
)

func newClassJSON(r, s string) *ClassJSON {
	return &ClassJSON{r, &s}
}

var noNullTimetablesResponse = TimetablesResponse{
	Timetables: TimetablesJSON{
		TimetableJSON{
			One:   newClassJSON("A", "1"),
			Two:   newClassJSON("B", "2"),
			Three: newClassJSON("C", "3"),
			Four:  newClassJSON("D", "4"),
			Five:  newClassJSON("E", "5"),
		},
		TimetableJSON{
			One:   newClassJSON("F", "6"),
			Two:   newClassJSON("G", "7"),
			Three: newClassJSON("H", "8"),
			Four:  newClassJSON("I", "9"),
			Five:  newClassJSON("J", "10"),
		},
		TimetableJSON{
			One:   newClassJSON("K", "11"),
			Two:   newClassJSON("L", "12"),
			Three: newClassJSON("M", "13"),
			Four:  newClassJSON("N", "14"),
			Five:  newClassJSON("O", "15"),
		},
		TimetableJSON{
			One:   newClassJSON("P", "16"),
			Two:   newClassJSON("Q", "17"),
			Three: newClassJSON("R", "18"),
			Four:  newClassJSON("S", "19"),
			Five:  newClassJSON("T", "20"),
		},
		TimetableJSON{
			One:   newClassJSON("U", "21"),
			Two:   newClassJSON("V", "22"),
			Three: newClassJSON("W", "23"),
			Four:  newClassJSON("X", "24"),
			Five:  newClassJSON("Y", "25"),
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

var hasNullClassTimetablesResponse = TimetablesResponse{
	Timetables: TimetablesJSON{
		TimetableJSON{
			One:   nil,
			Two:   newClassJSON("B", "2"),
			Three: newClassJSON("C", "3"),
			Four:  newClassJSON("D", "4"),
			Five:  newClassJSON("E", "5"),
		},
		TimetableJSON{
			One:   newClassJSON("F", "6"),
			Two:   nil,
			Three: newClassJSON("H", "8"),
			Four:  newClassJSON("I", "9"),
			Five:  newClassJSON("J", "10"),
		},
		TimetableJSON{
			One:   newClassJSON("K", "11"),
			Two:   newClassJSON("L", "12"),
			Three: nil,
			Four:  newClassJSON("N", "14"),
			Five:  newClassJSON("O", "15"),
		},
		TimetableJSON{
			One:   newClassJSON("P", "16"),
			Two:   newClassJSON("Q", "17"),
			Three: newClassJSON("R", "18"),
			Four:  nil,
			Five:  newClassJSON("T", "20"),
		},
		TimetableJSON{
			One:   newClassJSON("U", "21"),
			Two:   newClassJSON("V", "22"),
			Three: newClassJSON("W", "23"),
			Four:  newClassJSON("X", "24"),
			Five:  nil,
		},
	},
}

var hasNullClassTimetables = timetablesModel.NewTimetables(
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

var hasNullRoomTimetablesResponse = TimetablesResponse{
	Timetables: TimetablesJSON{
		TimetableJSON{
			One:   &ClassJSON{"A", nil},
			Two:   newClassJSON("B", "2"),
			Three: newClassJSON("C", "3"),
			Four:  newClassJSON("D", "4"),
			Five:  newClassJSON("E", "5"),
		},
		TimetableJSON{
			One:   newClassJSON("F", "6"),
			Two:   &ClassJSON{"G", nil},
			Three: newClassJSON("H", "8"),
			Four:  newClassJSON("I", "9"),
			Five:  newClassJSON("J", "10"),
		},
		TimetableJSON{
			One:   newClassJSON("K", "11"),
			Two:   newClassJSON("L", "12"),
			Three: &ClassJSON{"M", nil},
			Four:  newClassJSON("N", "14"),
			Five:  newClassJSON("O", "15"),
		},
		TimetableJSON{
			One:   newClassJSON("P", "16"),
			Two:   newClassJSON("Q", "17"),
			Three: newClassJSON("R", "18"),
			Four:  &ClassJSON{"S", nil},
			Five:  newClassJSON("T", "20"),
		},
		TimetableJSON{
			One:   newClassJSON("U", "21"),
			Two:   newClassJSON("V", "22"),
			Three: newClassJSON("W", "23"),
			Four:  newClassJSON("X", "24"),
			Five:  &ClassJSON{"Y", nil},
		},
	},
}

var hasNullRoomTimetables = timetablesModel.NewTimetables(
	timetablesModel.NewTimetable(
		timetablesModel.NoRoom("A"),
		timetablesModel.NewClass("B", "2"),
		timetablesModel.NewClass("C", "3"),
		timetablesModel.NewClass("D", "4"),
		timetablesModel.NewClass("E", "5"),
	),
	timetablesModel.NewTimetable(
		timetablesModel.NewClass("F", "6"),
		timetablesModel.NoRoom("G"),
		timetablesModel.NewClass("H", "8"),
		timetablesModel.NewClass("I", "9"),
		timetablesModel.NewClass("J", "10"),
	),
	timetablesModel.NewTimetable(
		timetablesModel.NewClass("K", "11"),
		timetablesModel.NewClass("L", "12"),
		timetablesModel.NoRoom("M"),
		timetablesModel.NewClass("N", "14"),
		timetablesModel.NewClass("O", "15"),
	),
	timetablesModel.NewTimetable(
		timetablesModel.NewClass("P", "16"),
		timetablesModel.NewClass("Q", "17"),
		timetablesModel.NewClass("R", "18"),
		timetablesModel.NoRoom("S"),
		timetablesModel.NewClass("T", "20"),
	),
	timetablesModel.NewTimetable(
		timetablesModel.NewClass("U", "21"),
		timetablesModel.NewClass("V", "22"),
		timetablesModel.NewClass("W", "23"),
		timetablesModel.NewClass("X", "24"),
		timetablesModel.NoRoom("Y"),
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
	Name:     "has null class",
	Input:    hasNullClassTimetablesResponse,
	Expected: hasNullClassTimetables,
}

var tr2t3 = TimetablesResponseToTimetables{
	Name:     "has null room",
	Input:    hasNullRoomTimetablesResponse,
	Expected: hasNullRoomTimetables,
}

func TestTimetablesResponseToTimetables(t *testing.T) {
	tcs := []TimetablesResponseToTimetables{
		tr2t1,
		tr2t2,
		tr2t3,
	}

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
	Name:     "has null class",
	Input:    hasNullClassTimetables,
	Expected: hasNullClassTimetablesResponse,
}

var t2tr3 = TimetablesToTimetablesResponse{
	Name:     "has null room",
	Input:    hasNullRoomTimetables,
	Expected: hasNullRoomTimetablesResponse,
}

func TestTimetablesToTimetablesResponse(t *testing.T) {
	tcs := []TimetablesToTimetablesResponse{
		t2tr1,
		t2tr2,
		t2tr3,
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			tt := toTimetablesResponse(tc.Input)
			if tt.toTimetables() != tc.Expected.toTimetables() {
				t.Errorf("Failed# expected: %v; got: %v\n", tc.Expected, tt)
			}
		})
	}
}

type TimetablesResponseValidation struct {
	Name     string
	Input    TimetablesResponse
	Expected bool
}

func (t TimetablesResponse) tooLongSubject() TimetablesResponse {
	tr := t.copy()
	tr.Timetables.Mon.One.Subject = strings.Repeat("A", 86)
	return tr
}

func (t TimetablesResponse) maxLengthSubject() TimetablesResponse {
	tr := t.copy()
	tr.Timetables.Mon.One.Subject = strings.Repeat("A", 85)
	return tr
}

func (t TimetablesResponse) tooLongRoom() TimetablesResponse {
	tr := t.copy()
	r := strings.Repeat("1", 86)
	tr.Timetables.Mon.One.Room = &r
	return tr
}

func (t TimetablesResponse) maxLengthRoom() TimetablesResponse {
	tr := t.copy()
	r := strings.Repeat("1", 85)
	tr.Timetables.Mon.One.Room = &r
	return tr
}

func TestValidates(t *testing.T) {
	tcs := []TimetablesResponseValidation{
		{
			Name:     "valid timetables have no null",
			Input:    noNullTimetablesResponse,
			Expected: true,
		},
		{
			Name:     "valid timetables have null classes",
			Input:    hasNullClassTimetablesResponse,
			Expected: true,
		},
		{
			Name:     "valid timetables have null rooms",
			Input:    hasNullRoomTimetablesResponse,
			Expected: true,
		},
		{
			Name:     "valid timetables have a max length subject",
			Input:    noNullTimetablesResponse.maxLengthSubject(),
			Expected: true,
		},
		{
			Name:     "invalid timetables have a too long subject",
			Input:    noNullTimetablesResponse.tooLongSubject(),
			Expected: false,
		},
		{
			Name:     "valid timetables have a max length room",
			Input:    noNullTimetablesResponse.maxLengthRoom(),
			Expected: true,
		},
		{
			Name:     "invalid timetables have a too long room",
			Input:    noNullTimetablesResponse.tooLongRoom(),
			Expected: false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			b, err := tc.Input.Validates()
			if err != nil {
				fmt.Println(tc.Input.Timetables.Mon.One.Subject)
				t.Errorf("unexpected error occured: %v", err)
			}
			if b != tc.Expected {
				t.Errorf("Failed# expected: %v; got: %v\n", tc.Expected, b)
			}
		})
	}
}

func (t TimetablesResponse) copy() TimetablesResponse {
	return TimetablesResponse{
		TimetablesJSON{
			Mon: t.Timetables.Mon.copy(),
			Tue: t.Timetables.Tue.copy(),
			Wed: t.Timetables.Wed.copy(),
			Thu: t.Timetables.Thu.copy(),
			Fri: t.Timetables.Fri.copy(),
		},
	}
}

func (t TimetableJSON) copy() TimetableJSON {
	_1 := t.One.copy()
	_2 := t.Two.copy()
	_3 := t.Three.copy()
	_4 := t.Four.copy()
	_5 := t.Five.copy()
	return TimetableJSON{
		One:   &_1,
		Two:   &_2,
		Three: &_3,
		Four:  &_4,
		Five:  &_5,
	}
}

func (t ClassJSON) copy() ClassJSON {
	var (
		s string
		r string
	)
	if t.Room != nil {
		s = t.Subject
		r = *t.Room
		return ClassJSON{s, &r}
	}

	s = t.Subject
	return ClassJSON{s, nil}
}
