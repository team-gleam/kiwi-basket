package timetables

import (
	"fmt"
	"strings"
	"testing"

	timetablesModel "github.com/team-gleam/kiwi-basket/server/src/domain/model/timetables"
)

func newClassJSON(s, r, m string) *ClassJSON {
	return &ClassJSON{s, &r, &m}
}

func newNoRoomClassJSON(s, m string) *ClassJSON {
	return &ClassJSON{s, nil, &m}
}

func newNoMemoClassJSON(s, r string) *ClassJSON {
	return &ClassJSON{s, &r, nil}
}

var noNullTimetablesResponse = TimetablesResponse{
	Timetables: TimetablesJSON{
		TimetableJSON{
			One:   newClassJSON("A", "1", "memo1"),
			Two:   newClassJSON("B", "2", "memo2"),
			Three: newClassJSON("C", "3", "memo3"),
			Four:  newClassJSON("D", "4", "memo4"),
			Five:  newClassJSON("E", "5", "memo5"),
		},
		TimetableJSON{
			One:   newClassJSON("F", "6", "memo6"),
			Two:   newClassJSON("G", "7", "memo7"),
			Three: newClassJSON("H", "8", "memo8"),
			Four:  newClassJSON("I", "9", "memo9"),
			Five:  newClassJSON("J", "10", "memo10"),
		},
		TimetableJSON{
			One:   newClassJSON("K", "11", "memo11"),
			Two:   newClassJSON("L", "12", "memo12"),
			Three: newClassJSON("M", "13", "memo13"),
			Four:  newClassJSON("N", "14", "memo14"),
			Five:  newClassJSON("O", "15", "memo15"),
		},
		TimetableJSON{
			One:   newClassJSON("P", "16", "memo16"),
			Two:   newClassJSON("Q", "17", "memo17"),
			Three: newClassJSON("R", "18", "memo18"),
			Four:  newClassJSON("S", "19", "memo19"),
			Five:  newClassJSON("T", "20", "memo20"),
		},
		TimetableJSON{
			One:   newClassJSON("U", "21", "memo21"),
			Two:   newClassJSON("V", "22", "memo22"),
			Three: newClassJSON("W", "23", "memo23"),
			Four:  newClassJSON("X", "24", "memo24"),
			Five:  newClassJSON("Y", "25", "memo25"),
		},
	},
}

var noNullTimetables = timetablesModel.NewTimetables(
	timetablesModel.NewTimetable(
		timetablesModel.NewClass("A", "1", "memo1"),
		timetablesModel.NewClass("B", "2", "memo2"),
		timetablesModel.NewClass("C", "3", "memo3"),
		timetablesModel.NewClass("D", "4", "memo4"),
		timetablesModel.NewClass("E", "5", "memo5"),
	),
	timetablesModel.NewTimetable(
		timetablesModel.NewClass("F", "6", "memo6"),
		timetablesModel.NewClass("G", "7", "memo7"),
		timetablesModel.NewClass("H", "8", "memo8"),
		timetablesModel.NewClass("I", "9", "memo9"),
		timetablesModel.NewClass("J", "10", "memo10"),
	),
	timetablesModel.NewTimetable(
		timetablesModel.NewClass("K", "11", "memo11"),
		timetablesModel.NewClass("L", "12", "memo12"),
		timetablesModel.NewClass("M", "13", "memo13"),
		timetablesModel.NewClass("N", "14", "memo14"),
		timetablesModel.NewClass("O", "15", "memo15"),
	),
	timetablesModel.NewTimetable(
		timetablesModel.NewClass("P", "16", "memo16"),
		timetablesModel.NewClass("Q", "17", "memo17"),
		timetablesModel.NewClass("R", "18", "memo18"),
		timetablesModel.NewClass("S", "19", "memo19"),
		timetablesModel.NewClass("T", "20", "memo20"),
	),
	timetablesModel.NewTimetable(
		timetablesModel.NewClass("U", "21", "memo21"),
		timetablesModel.NewClass("V", "22", "memo22"),
		timetablesModel.NewClass("W", "23", "memo23"),
		timetablesModel.NewClass("X", "24", "memo24"),
		timetablesModel.NewClass("Y", "25", "memo25"),
	),
)

var hasNullClassTimetablesResponse = TimetablesResponse{
	Timetables: TimetablesJSON{
		TimetableJSON{
			One:   nil,
			Two:   newClassJSON("B", "2", "memo2"),
			Three: newClassJSON("C", "3", "memo3"),
			Four:  newClassJSON("D", "4", "memo4"),
			Five:  newClassJSON("E", "5", "memo5"),
		},
		TimetableJSON{
			One:   newClassJSON("F", "6", "memo6"),
			Two:   nil,
			Three: newClassJSON("H", "8", "memo8"),
			Four:  newClassJSON("I", "9", "memo9"),
			Five:  newClassJSON("J", "10", "memo10"),
		},
		TimetableJSON{
			One:   newClassJSON("K", "11", "memo11"),
			Two:   newClassJSON("L", "12", "memo12"),
			Three: nil,
			Four:  newClassJSON("N", "14", "memo14"),
			Five:  newClassJSON("O", "15", "memo15"),
		},
		TimetableJSON{
			One:   newClassJSON("P", "16", "memo16"),
			Two:   newClassJSON("Q", "17", "memo17"),
			Three: newClassJSON("R", "18", "memo18"),
			Four:  nil,
			Five:  newClassJSON("T", "20", "memo20"),
		},
		TimetableJSON{
			One:   newClassJSON("U", "21", "memo21"),
			Two:   newClassJSON("V", "22", "memo22"),
			Three: newClassJSON("W", "23", "memo23"),
			Four:  newClassJSON("X", "24", "memo24"),
			Five:  nil,
		},
	},
}

var hasNullClassTimetables = timetablesModel.NewTimetables(
	timetablesModel.NewTimetable(
		timetablesModel.NoClass(),
		timetablesModel.NewClass("B", "2", "memo2"),
		timetablesModel.NewClass("C", "3", "memo3"),
		timetablesModel.NewClass("D", "4", "memo4"),
		timetablesModel.NewClass("E", "5", "memo5"),
	),
	timetablesModel.NewTimetable(
		timetablesModel.NewClass("F", "6", "memo6"),
		timetablesModel.NoClass(),
		timetablesModel.NewClass("H", "8", "memo8"),
		timetablesModel.NewClass("I", "9", "memo9"),
		timetablesModel.NewClass("J", "10", "memo10"),
	),
	timetablesModel.NewTimetable(
		timetablesModel.NewClass("K", "11", "memo11"),
		timetablesModel.NewClass("L", "12", "memo12"),
		timetablesModel.NoClass(),
		timetablesModel.NewClass("N", "14", "memo14"),
		timetablesModel.NewClass("O", "15", "memo15"),
	),
	timetablesModel.NewTimetable(
		timetablesModel.NewClass("P", "16", "memo16"),
		timetablesModel.NewClass("Q", "17", "memo17"),
		timetablesModel.NewClass("R", "18", "memo18"),
		timetablesModel.NoClass(),
		timetablesModel.NewClass("T", "20", "memo20"),
	),
	timetablesModel.NewTimetable(
		timetablesModel.NewClass("U", "21", "memo21"),
		timetablesModel.NewClass("V", "22", "memo22"),
		timetablesModel.NewClass("W", "23", "memo23"),
		timetablesModel.NewClass("X", "24", "memo24"),
		timetablesModel.NoClass(),
	),
)

var hasNullRoomTimetablesResponse = TimetablesResponse{
	Timetables: TimetablesJSON{
		TimetableJSON{
			One:   newNoRoomClassJSON("A", "memo1"),
			Two:   newClassJSON("B", "2", "memo2"),
			Three: newClassJSON("C", "3", "memo3"),
			Four:  newClassJSON("D", "4", "memo4"),
			Five:  newClassJSON("E", "5", "memo5"),
		},
		TimetableJSON{
			One:   newClassJSON("F", "6", "memo6"),
			Two:   newNoRoomClassJSON("G", "memo7"),
			Three: newClassJSON("H", "8", "memo8"),
			Four:  newClassJSON("I", "9", "memo9"),
			Five:  newClassJSON("J", "10", "memo10"),
		},
		TimetableJSON{
			One:   newClassJSON("K", "11", "memo11"),
			Two:   newClassJSON("L", "12", "memo12"),
			Three: newNoRoomClassJSON("M", "memo13"),
			Four:  newClassJSON("N", "14", "memo14"),
			Five:  newClassJSON("O", "15", "memo15"),
		},
		TimetableJSON{
			One:   newClassJSON("P", "16", "memo16"),
			Two:   newClassJSON("Q", "17", "memo17"),
			Three: newClassJSON("R", "18", "memo18"),
			Four:  newNoRoomClassJSON("S", "memo19"),
			Five:  newClassJSON("T", "20", "memo20"),
		},
		TimetableJSON{
			One:   newClassJSON("U", "21", "memo21"),
			Two:   newClassJSON("V", "22", "memo22"),
			Three: newClassJSON("W", "23", "memo23"),
			Four:  newClassJSON("X", "24", "memo24"),
			Five:  newNoRoomClassJSON("Y", "memo25"),
		},
	},
}

var hasNullRoomTimetables = timetablesModel.NewTimetables(
	timetablesModel.NewTimetable(
		timetablesModel.NoRoom("A", "memo1"),
		timetablesModel.NewClass("B", "2", "memo2"),
		timetablesModel.NewClass("C", "3", "memo3"),
		timetablesModel.NewClass("D", "4", "memo4"),
		timetablesModel.NewClass("E", "5", "memo5"),
	),
	timetablesModel.NewTimetable(
		timetablesModel.NewClass("F", "6", "memo6"),
		timetablesModel.NoRoom("G", "memo7"),
		timetablesModel.NewClass("H", "8", "memo8"),
		timetablesModel.NewClass("I", "9", "memo9"),
		timetablesModel.NewClass("J", "10", "memo10"),
	),
	timetablesModel.NewTimetable(
		timetablesModel.NewClass("K", "11", "memo11"),
		timetablesModel.NewClass("L", "12", "memo12"),
		timetablesModel.NoRoom("M", "memo13"),
		timetablesModel.NewClass("N", "14", "memo14"),
		timetablesModel.NewClass("O", "15", "memo15"),
	),
	timetablesModel.NewTimetable(
		timetablesModel.NewClass("P", "16", "memo16"),
		timetablesModel.NewClass("Q", "17", "memo17"),
		timetablesModel.NewClass("R", "18", "memo18"),
		timetablesModel.NoRoom("S", "memo19"),
		timetablesModel.NewClass("T", "20", "memo20"),
	),
	timetablesModel.NewTimetable(
		timetablesModel.NewClass("U", "21", "memo21"),
		timetablesModel.NewClass("V", "22", "memo22"),
		timetablesModel.NewClass("W", "23", "memo23"),
		timetablesModel.NewClass("X", "24", "memo24"),
		timetablesModel.NoRoom("Y", "memo25"),
	),
)

var hasNullMemoTimetablesResponse = TimetablesResponse{
	Timetables: TimetablesJSON{
		TimetableJSON{
			One:   newNoMemoClassJSON("A", "1"),
			Two:   newClassJSON("B", "2", "memo2"),
			Three: newClassJSON("C", "3", "memo3"),
			Four:  newClassJSON("D", "4", "memo4"),
			Five:  newClassJSON("E", "5", "memo5"),
		},
		TimetableJSON{
			One:   newClassJSON("F", "6", "memo6"),
			Two:   newNoMemoClassJSON("G", "7"),
			Three: newClassJSON("H", "8", "memo8"),
			Four:  newClassJSON("I", "9", "memo9"),
			Five:  newClassJSON("J", "10", "memo10"),
		},
		TimetableJSON{
			One:   newClassJSON("K", "11", "memo11"),
			Two:   newClassJSON("L", "12", "memo12"),
			Three: newNoMemoClassJSON("M", "13"),
			Four:  newClassJSON("N", "14", "memo14"),
			Five:  newClassJSON("O", "15", "memo15"),
		},
		TimetableJSON{
			One:   newClassJSON("P", "16", "memo16"),
			Two:   newClassJSON("Q", "17", "memo17"),
			Three: newClassJSON("R", "18", "memo18"),
			Four:  newNoMemoClassJSON("S", "19"),
			Five:  newClassJSON("T", "20", "memo20"),
		},
		TimetableJSON{
			One:   newClassJSON("U", "21", "memo21"),
			Two:   newClassJSON("V", "22", "memo22"),
			Three: newClassJSON("W", "23", "memo23"),
			Four:  newClassJSON("X", "24", "memo24"),
			Five:  newNoMemoClassJSON("Y", "25"),
		},
	},
}

var hasNullMemoTimetables = timetablesModel.NewTimetables(
	timetablesModel.NewTimetable(
		timetablesModel.NewClass("A", "1", ""),
		timetablesModel.NewClass("B", "2", "memo2"),
		timetablesModel.NewClass("C", "3", "memo3"),
		timetablesModel.NewClass("D", "4", "memo4"),
		timetablesModel.NewClass("E", "5", "memo5"),
	),
	timetablesModel.NewTimetable(
		timetablesModel.NewClass("F", "6", "memo6"),
		timetablesModel.NewClass("G", "7", ""),
		timetablesModel.NewClass("H", "8", "memo8"),
		timetablesModel.NewClass("I", "9", "memo9"),
		timetablesModel.NewClass("J", "10", "memo10"),
	),
	timetablesModel.NewTimetable(
		timetablesModel.NewClass("K", "11", "memo11"),
		timetablesModel.NewClass("L", "12", "memo12"),
		timetablesModel.NewClass("M", "13", ""),
		timetablesModel.NewClass("N", "14", "memo14"),
		timetablesModel.NewClass("O", "15", "memo15"),
	),
	timetablesModel.NewTimetable(
		timetablesModel.NewClass("P", "16", "memo16"),
		timetablesModel.NewClass("Q", "17", "memo17"),
		timetablesModel.NewClass("R", "18", "memo18"),
		timetablesModel.NewClass("S", "19", ""),
		timetablesModel.NewClass("T", "20", "memo20"),
	),
	timetablesModel.NewTimetable(
		timetablesModel.NewClass("U", "21", "memo21"),
		timetablesModel.NewClass("V", "22", "memo22"),
		timetablesModel.NewClass("W", "23", "memo23"),
		timetablesModel.NewClass("X", "24", "memo24"),
		timetablesModel.NewClass("Y", "25", ""),
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

var tr2t4 = TimetablesResponseToTimetables{
	Name:     "has null memo",
	Input:    hasNullMemoTimetablesResponse,
	Expected: hasNullMemoTimetables,
}

func TestTimetablesResponseToTimetables(t *testing.T) {
	tcs := []TimetablesResponseToTimetables{
		tr2t1,
		tr2t2,
		tr2t3,
		tr2t4,
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			tt := tc.Input.toTimetables()
			if tt != tc.Expected {
				t.Fatalf("expected: %v; got: %v\n", tc.Expected, tt)
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

var t2tr4 = TimetablesToTimetablesResponse{
	Name:     "has null memo",
	Input:    hasNullMemoTimetables,
	Expected: hasNullMemoTimetablesResponse,
}

func TestTimetablesToTimetablesResponse(t *testing.T) {
	tcs := []TimetablesToTimetablesResponse{
		t2tr1,
		t2tr2,
		t2tr3,
		t2tr4,
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			tt := toTimetablesResponse(tc.Input)
			if tt.toTimetables() != tc.Expected.toTimetables() {
				t.Fatalf("expected: %v; got: %v\n", tc.Expected, tt)
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
				t.Fatalf("unexpected error occured: %v", err)
			}
			if b != tc.Expected {
				t.Fatalf("expected: %v; got: %v\n", tc.Expected, b)
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
		m string
	)
	s = t.Subject

	if t.Room != nil {
		r = *t.Room
		if t.Memo != nil {
			m = *t.Memo
			return ClassJSON{s, &r, &m}
		}
		return ClassJSON{s, &r, nil}
	}

	if t.Memo != nil {
		m = *t.Memo
		return ClassJSON{s, nil, &m}
	}
	return ClassJSON{s, nil, nil}
}
