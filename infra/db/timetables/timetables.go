package timetables

import (
	timetablesModel "github.com/team-gleam/kiwi-basket/domain/model/timetables"
	"github.com/team-gleam/kiwi-basket/domain/model/user/username"
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

const (
	NoClass = "no class"
	Mon     = "mon"
	Tue     = "tue"
	Wed     = "wed"
	Thu     = "thu"
	Fri     = "fri"
)

func (r *TimetablesRepository) Create(u username.Username, t timetablesModel.Timetables) error {
	mon, err := r.createTimetable(Mon, t.Mon())
	tue, err := r.createTimetable(Tue, t.Tue())
	wed, err := r.createTimetable(Wed, t.Wed())
	thu, err := r.createTimetable(Thu, t.Thu())
	fri, err := r.createTimetable(Fri, t.Fri())
	if err != nil {
		return err
	}

	return r.dbHandler.Db.Create(NewTimetablesDB(u.Name(), mon, tue, wed, thu, fri)).Error
}

func (r *TimetablesRepository) createTimetable(day string, timetable timetablesModel.Timetable) (uint, error) {
	_1, err := r.createClass(timetable.First())
	_2, err := r.createClass(timetable.Second())
	_3, err := r.createClass(timetable.Third())
	_4, err := r.createClass(timetable.Fourth())
	_5, err := r.createClass(timetable.Fifth())
	if err != nil {
		return 0, err
	}

	t := NewTimetableDB(day, _1, _2, _3, _4, _5)
	err = r.dbHandler.Db.Create(&t).Error
	return t.ID, err
}

func (r *TimetablesRepository) createClass(class timetablesModel.Class) (*uint, error) {
	if class.IsNoClass() {
		return nil, nil
	}

	c := NewClassDB(class.Subject(), class.Room())
	err := r.dbHandler.Db.Create(&c).Error
	return &c.ID, err
}

func (r *TimetablesRepository) Delete(u username.Username) error {
	return r.dbHandler.Db.Delete(TimetablesDB{Username: u.Name()}).Error
}

func (r *TimetablesRepository) Exists(u username.Username) (bool, error) {
	t := TimetablesDB{}
	err := r.dbHandler.Db.Where("username = ?", u.Name()).First(&t).Error
	if err != nil {
		return false, err
	}

	return t != TimetablesDB{}, nil
}
