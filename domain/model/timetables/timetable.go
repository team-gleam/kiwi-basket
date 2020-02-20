package timetables

type Timetable struct {
	_1 Class
	_2 Class
	_3 Class
	_4 Class
	_5 Class
}

func NewTimetable(_1, _2, _3, _4, _5 Class) Timetable {
	return Timetable{_1, _2, _3, _4, _5}
}
