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

func (t Timetables) Mon() Timetable {
	return t.mon
}

func (t Timetables) Tue() Timetable {
	return t.tue
}

func (t Timetables) Wed() Timetable {
	return t.wed
}

func (t Timetables) Thu() Timetable {
	return t.thu
}

func (t Timetables) Fri() Timetable {
	return t.fri
}
