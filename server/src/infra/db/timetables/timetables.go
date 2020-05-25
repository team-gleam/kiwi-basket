package timetables

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	timetablesModel "github.com/team-gleam/kiwi-basket/domain/model/timetables"
	"github.com/team-gleam/kiwi-basket/domain/model/user/username"
	timetablesRepository "github.com/team-gleam/kiwi-basket/domain/repository/timetables"
	"github.com/team-gleam/kiwi-basket/infra/db/handler"
)

type TimetablesRepository struct {
	dbHandler *handler.DbHandler
}

func NewTimetablesRepository(h *handler.DbHandler) timetablesRepository.ITimetablesRepository {
	h.Db.AutoMigrate(
		Timetables{},
		Timetable{},
		Class{},
	)
	return &TimetablesRepository{h}
}

type Timetables struct {
	Username                string `gorm:"primary_key"`
	Mon, Tue, Wed, Thu, Fri uint
}

func NewTimetables(u string, mon, tue, wed, thu, fri uint) Timetables {
	return Timetables{
		Username: u,
		Mon:      mon,
		Tue:      tue,
		Wed:      wed,
		Thu:      thu,
		Fri:      fri,
	}
}

func (t Timetables) TableName() string {
	return "timetables"
}

type Timetable struct {
	ID                          uint `gorm:"primary_key;auto_increment"`
	Day                         string
	One, Two, Three, Four, Five *uint
}

func NewTimetable(d string, _1, _2, _3, _4, _5 *uint) Timetable {
	return Timetable{
		Day:   d,
		One:   _1,
		Two:   _2,
		Three: _3,
		Four:  _4,
		Five:  _5,
	}
}

func (t Timetable) TableName() string {
	return "timetable"
}

type Class struct {
	ID      uint `gorm:"primary_key;auto_increment"`
	Subject string
	Room    sql.NullString
}

func NewClass(s, r string) Class {
	return Class{
		Subject: s,
		Room: sql.NullString{
			String: r,
			Valid:  true,
		},
	}
}

func NewNoRoomClass(s string) Class {
	return Class{
		Subject: s,
		Room: sql.NullString{
			Valid: false,
		},
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
	if err != nil {
		return err
	}
	tue, err := r.createTimetable(Tue, t.Tue())
	if err != nil {
		return err
	}
	wed, err := r.createTimetable(Wed, t.Wed())
	if err != nil {
		return err
	}
	thu, err := r.createTimetable(Thu, t.Thu())
	if err != nil {
		return err
	}
	fri, err := r.createTimetable(Fri, t.Fri())
	if err != nil {
		return err
	}

	return r.dbHandler.Db.Create(NewTimetables(u.Name(), mon, tue, wed, thu, fri)).Error
}

func (r *TimetablesRepository) createTimetable(day string, timetable timetablesModel.Timetable) (uint, error) {
	_1, err := r.createClass(timetable.First())
	if err != nil {
		return 0, err
	}
	_2, err := r.createClass(timetable.Second())
	if err != nil {
		return 0, err
	}
	_3, err := r.createClass(timetable.Third())
	if err != nil {
		return 0, err
	}
	_4, err := r.createClass(timetable.Fourth())
	if err != nil {
		return 0, err
	}
	_5, err := r.createClass(timetable.Fifth())
	if err != nil {
		return 0, err
	}

	t := NewTimetable(day, _1, _2, _3, _4, _5)
	err = r.dbHandler.Db.Create(&t).Error
	return t.ID, err
}

func (r *TimetablesRepository) createClass(class timetablesModel.Class) (*uint, error) {
	if class.IsNoClass() {
		return nil, nil
	}

	var c Class
	if !class.IsNoRoom() {
		c = NewClass(class.Subject(), class.Room())
	} else {
		c = NewNoRoomClass(class.Subject())
	}

	err := r.dbHandler.Db.Create(&c).Error
	return &c.ID, err
}

func (r *TimetablesRepository) Delete(u username.Username) error {
	ts := new(Timetables)
	err := r.dbHandler.Db.Where("username = ?", u.Name()).Take(ts).Error
	if err != nil {
		return err
	}

	err = r.deleteTimetable(ts.Mon)
	if err != nil {
		return err
	}
	err = r.deleteTimetable(ts.Tue)
	if err != nil {
		return err
	}
	err = r.deleteTimetable(ts.Wed)
	if err != nil {
		return err
	}
	err = r.deleteTimetable(ts.Thu)
	if err != nil {
		return err
	}
	err = r.deleteTimetable(ts.Fri)
	if err != nil {
		return err
	}

	return r.dbHandler.Db.Where("username = ?", u.Name()).Delete(Timetables{}).Error
}

func (r *TimetablesRepository) deleteTimetable(id uint) error {
	td := new(Timetable)
	err := r.dbHandler.Db.Where("id = ?", id).Take(td).Error
	if err != nil {
		return err
	}

	err = r.deleteClass(td.One)
	if err != nil {
		return err
	}
	err = r.deleteClass(td.Two)
	if err != nil {
		return err
	}
	err = r.deleteClass(td.Three)
	if err != nil {
		return err
	}
	err = r.deleteClass(td.Four)
	if err != nil {
		return err
	}
	err = r.deleteClass(td.Five)
	if err != nil {
		return err
	}

	return r.dbHandler.Db.Where("id = ?", id).Delete(Timetable{}).Error
}

func (r *TimetablesRepository) deleteClass(id *uint) error {
	if id != nil {
		return r.dbHandler.Db.Where("id = ?", *id).Delete(Class{}).Error
	}

	return nil
}

func (r *TimetablesRepository) Exists(u username.Username) (bool, error) {
	t := Timetables{}
	err := r.dbHandler.Db.Where("username = ?", u.Name()).Take(&t).Error
	if gorm.IsRecordNotFoundError(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return t != Timetables{}, nil
}

func (r *TimetablesRepository) Get(u username.Username) (timetablesModel.Timetables, error) {
	ts := Timetables{}
	err := r.dbHandler.Db.Where("username = ?", u.Name()).Take(&ts).Error
	if err != nil {
		return timetablesModel.Timetables{}, err
	}

	mon, err := r.getTimetable(ts.Mon)
	if err != nil {
		return timetablesModel.Timetables{}, err
	}
	tue, err := r.getTimetable(ts.Tue)
	if err != nil {
		return timetablesModel.Timetables{}, err
	}
	wed, err := r.getTimetable(ts.Wed)
	if err != nil {
		return timetablesModel.Timetables{}, err
	}
	thu, err := r.getTimetable(ts.Thu)
	if err != nil {
		return timetablesModel.Timetables{}, err
	}
	fri, err := r.getTimetable(ts.Fri)
	if err != nil {
		return timetablesModel.Timetables{}, err
	}

	return timetablesModel.NewTimetables(mon, tue, wed, thu, fri), nil
}

func (r *TimetablesRepository) getTimetable(id uint) (timetablesModel.Timetable, error) {
	t := Timetable{}
	err := r.dbHandler.Db.Where("id = ?", id).Take(&t).Error
	if err != nil {
		return timetablesModel.Timetable{}, err
	}

	_1, err := r.getClass(t.One)
	if err != nil {
		return timetablesModel.Timetable{}, err
	}
	_2, err := r.getClass(t.Two)
	if err != nil {
		return timetablesModel.Timetable{}, err
	}
	_3, err := r.getClass(t.Three)
	if err != nil {
		return timetablesModel.Timetable{}, err
	}
	_4, err := r.getClass(t.Four)
	if err != nil {
		return timetablesModel.Timetable{}, err
	}
	_5, err := r.getClass(t.Five)
	if err != nil {
		return timetablesModel.Timetable{}, err
	}

	return timetablesModel.NewTimetable(_1, _2, _3, _4, _5), nil
}

func (r *TimetablesRepository) getClass(id *uint) (timetablesModel.Class, error) {
	if id == nil {
		return timetablesModel.NoClass(), nil
	}

	c := Class{}
	err := r.dbHandler.Db.Where("id = ?", id).Take(&c).Error
	if err != nil {
		return timetablesModel.Class{}, err
	}

	if !c.Room.Valid {
		return timetablesModel.NoRoom(c.Subject), nil
	}

	return timetablesModel.NewClass(c.Subject, c.Room.String), nil
}
