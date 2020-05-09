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
		TimetablesDB{},
		TimetableDB{},
		ClassDB{},
	)
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
	One, Two, Three, Four, Five *uint
}

func NewTimetableDB(d string, _1, _2, _3, _4, _5 *uint) TimetableDB {
	return TimetableDB{
		Day:   d,
		One:   _1,
		Two:   _2,
		Three: _3,
		Four:  _4,
		Five:  _5,
	}
}

type ClassDB struct {
	ID      uint `gorm:"primary_key;auto_increment"`
	Subject string
	Room    sql.NullString
}

func NewClassDB(s, r string) ClassDB {
	return ClassDB{
		Subject: s,
		Room: sql.NullString{
			String: r,
			Valid:  true,
		},
	}
}

func NewNoRoomClassDB(s string) ClassDB {
	return ClassDB{
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

	return r.dbHandler.Db.Create(NewTimetablesDB(u.Name(), mon, tue, wed, thu, fri)).Error
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

	t := NewTimetableDB(day, _1, _2, _3, _4, _5)
	err = r.dbHandler.Db.Create(&t).Error
	return t.ID, err
}

func (r *TimetablesRepository) createClass(class timetablesModel.Class) (*uint, error) {
	if class.IsNoClass() {
		return nil, nil
	}

	var c ClassDB
	if !class.IsNoRoom() {
		c = NewClassDB(class.Subject(), class.Room())
	} else {
		c = NewNoRoomClassDB(class.Subject())
	}

	err := r.dbHandler.Db.Create(&c).Error
	return &c.ID, err
}

func (r *TimetablesRepository) Delete(u username.Username) error {
	ts := new(TimetablesDB)
	err := r.dbHandler.Db.Where("username = ?", u.Name).Take(&ts).Error
	if err != nil {
		return err
	}

	err = r.deleteTimetable(TimetableDB{ID: ts.Mon})
	if err != nil {
		return err
	}
	err = r.deleteTimetable(TimetableDB{ID: ts.Tue})
	if err != nil {
		return err
	}
	err = r.deleteTimetable(TimetableDB{ID: ts.Wed})
	if err != nil {
		return err
	}
	err = r.deleteTimetable(TimetableDB{ID: ts.Thu})
	if err != nil {
		return err
	}
	err = r.deleteTimetable(TimetableDB{ID: ts.Fri})
	if err != nil {
		return err
	}

	return r.dbHandler.Db.Delete(TimetablesDB{Username: u.Name()}).Error
}

func (r *TimetablesRepository) deleteTimetable(t TimetableDB) error {
	err := r.deleteClass(ClassDB{ID: *t.One})
	if err != nil {
		return err
	}
	err = r.deleteClass(ClassDB{ID: *t.Two})
	if err != nil {
		return err
	}
	err = r.deleteClass(ClassDB{ID: *t.Three})
	if err != nil {
		return err
	}
	err = r.deleteClass(ClassDB{ID: *t.Four})
	if err != nil {
		return err
	}
	return r.deleteClass(ClassDB{ID: *t.Five})
}

func (r *TimetablesRepository) deleteClass(c ClassDB) error {
	return r.dbHandler.Db.Delete(c).Error
}

func (r *TimetablesRepository) Exists(u username.Username) (bool, error) {
	t := TimetablesDB{}
	err := r.dbHandler.Db.Where("username = ?", u.Name()).Take(&t).Error
	if gorm.IsRecordNotFoundError(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return t != TimetablesDB{}, nil
}

func (r *TimetablesRepository) Get(u username.Username) (timetablesModel.Timetables, error) {
	ts := TimetablesDB{}
	err := r.dbHandler.Db.Where("username = ?", u.Name).Take(&ts).Error
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
	t := TimetableDB{}
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

	c := ClassDB{}
	err := r.dbHandler.Db.Where("id = ?", id).Take(&c).Error
	if err != nil {
		return timetablesModel.Class{}, err
	}

	if !c.Room.Valid {
		return timetablesModel.NoRoom(c.Subject), nil
	}

	return timetablesModel.NewClass(c.Subject, c.Room.String), nil
}
