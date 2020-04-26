package timetables

import (
	timetablesRepository "github.com/team-gleam/kiwi-basket/domain/repository/timetables"
	"github.com/team-gleam/kiwi-basket/infra/db/handler"
)

type TimetablesRepository struct {
	dbHandler *handler.DbHandler
}

func NewTimetablesRepository(h *handler.DbHandler) timetablesRepository.ITimetablesRepository {
	return &TimetablesRepository{h}
}

type TimetablesDB struct {
	Username                string `gorm:"primary_key"`
	Mon, Tue, Wed, Thu, Fri uint
}

func NewTimetablesDB(u string, mon, tue, wed, thu, fri uint) TimetablesDB {
	return TimetablesDB{
		Username: u,
		Mon:      mon,
		Tue:      tue,
		Wed:      wed,
		Thu:      thu,
		Fri:      fri,
	}
}

type TimetableDB struct {
	ID                          uint `gorm:"primary_key;auto_increment"`
	Day                         string
	One, Twe, Three, Four, Five *uint
}

func NewTimetableDB(d string, _1, _2, _3, _4, _5 *uint) TimetableDB {
	return TimetableDB{
		Day:   d,
		One:   _1,
		Twe:   _2,
		Three: _3,
		Four:  _4,
		Five:  _5,
	}
}

type ClassDB struct {
	ID            uint `gorm:"primary_key;auto_increment"`
	Subject, Room string
}

func NewClassDB(s, r string) ClassDB {
	return ClassDB{
		Subject: s,
		Room:    r,
	}
}
