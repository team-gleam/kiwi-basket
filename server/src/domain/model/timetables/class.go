package timetables

type Class struct {
	subject string
	room    string
	noRoom  bool
	noClass bool
}

func NewClass(s, r string) Class {
	return Class{
		subject: s,
		room:    r,
		noRoom:  false,
		noClass: false,
	}
}

func NoClass() Class {
	return Class{noClass: true}
}

func NoRoom(s string) Class {
	return Class{
		subject: s,
		noRoom:  true,
	}
}

func (c Class) IsNoClass() bool {
	return c.noClass
}

func (c Class) IsNoRoom() bool {
	return c.noRoom
}

func (c Class) Subject() string {
	return c.subject
}

func (c Class) Room() string {
	return c.room
}
