package timetables

type Timetables struct {
	mon Timetable
	tue Timetable
	wed Timetable
	thu Timetable
	fri Timetable
}

func NewTimetables(mon, tue, wed, thu, fri Timetable) Timetables {
	return NewTimetables(mon, tue, wed, thu, fri)
}
