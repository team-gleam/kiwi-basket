package login

import (
	"fmt"

	"github.com/jinzhu/gorm"
	loginModel "github.com/team-gleam/kiwi-basket/server/src/domain/model/user/login"
	"github.com/team-gleam/kiwi-basket/server/src/domain/model/user/username"
	loginRepository "github.com/team-gleam/kiwi-basket/server/src/domain/repository/user/login"
	"github.com/team-gleam/kiwi-basket/server/src/infra/db/handler"
)

type LoginRepository struct {
	dbHandler *handler.DbHandler
}

func NewLoginRepository(h *handler.DbHandler) loginRepository.ILoginRepository {
	h.Db.AutoMigrate(Login{})
	return &LoginRepository{h}
}

type Login struct {
	Username string `gorm:"primary_key"`
	Password string
}

func toRecord(l loginModel.Login) Login {
	return Login{l.Username().Name(), l.HashedPassword()}
}

func fromRecord(l Login) (loginModel.Login, error) {
	u, err := username.NewUsername(l.Username)
	return loginModel.NewLogin(u, l.Password), err
}

func (r *LoginRepository) Create(l loginModel.Login) error {
	login := toRecord(l)
	return r.dbHandler.Db.Create(&login).Error
}

func (r *LoginRepository) Delete(l loginModel.Login) error {
	login := toRecord(l)
	return r.dbHandler.Db.Delete(login).Error
}

func (r *LoginRepository) Exists(u username.Username) (bool, error) {
	l := new(Login)
	err := r.dbHandler.Db.Where("username = ?", u.Name()).Take(l).Error
	if gorm.IsRecordNotFoundError(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return l.Username != "", nil
}

func (r *LoginRepository) Get(u username.Username) (loginModel.Login, error) {
	login := new(Login)
	err := r.dbHandler.Db.Where("username = ?", u.Name()).Take(login).Error
	if err != nil {
		return loginModel.Login{}, err
	}

	l, err := fromRecord(*login)
	if err != nil && err.Error() == username.InvalidUsername {
		return loginModel.Login{}, fmt.Errorf("user not found")
	}

	return l, nil
}
