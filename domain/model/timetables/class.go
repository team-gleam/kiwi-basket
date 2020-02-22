package timetables

type Class struct {
	subject string
	room    string
	noClass bool
}

func NewClass(s, r string) Class {
	return Class{s, r, false}
}

func NoClass() Class {
	return Class{noClass: true}
}

func (c Class) IsNoClass() bool {
	return c.noClass
}

func (c Class) Subject() string {
	return c.subject
}

func (c Class) Room() string {
	return c.room
}
