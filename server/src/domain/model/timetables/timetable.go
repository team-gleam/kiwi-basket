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

func (t Timetable) First() Class {
	return t._1
}

func (t Timetable) Second() Class {
	return t._2
}

func (t Timetable) Third() Class {
	return t._3
}

func (t Timetable) Fourth() Class {
	return t._4
}

func (t Timetable) Fifth() Class {
	return t._5
}
