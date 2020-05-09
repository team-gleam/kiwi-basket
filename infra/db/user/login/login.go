package login

import (
	"fmt"

	loginModel "github.com/team-gleam/kiwi-basket/domain/model/user/login"
	"github.com/team-gleam/kiwi-basket/domain/model/user/username"
	loginRepository "github.com/team-gleam/kiwi-basket/domain/repository/user/login"
	"github.com/team-gleam/kiwi-basket/infra/db/handler"
)

type LoginRepository struct {
	dbHandler *handler.DbHandler
}

func NewLoginRepository(h *handler.DbHandler) loginRepository.ILoginRepository {
	h.Db.AutoMigrate(LoginDB{})
	return &LoginRepository{h}
}

type LoginDB struct {
	Username string `gorm:"primary_key"`
	Password string
}

func transformLoginForDB(l loginModel.Login) LoginDB {
	return LoginDB{l.Username().Name(), l.HashedPassword()}
}

func toLogin(l LoginDB) (loginModel.Login, error) {
	u, err := username.NewUsername(l.Username)
	return loginModel.NewLogin(u, l.Password), err
}

func (r *LoginRepository) Create(l loginModel.Login) error {
	login := transformLoginForDB(l)
	return r.dbHandler.Db.Create(login).Error
}

func (r *LoginRepository) Delete(l loginModel.Login) error {
	login := transformLoginForDB(l)
	return r.dbHandler.Db.Delete(login).Error
}

func (r *LoginRepository) Exists(u username.Username) (bool, error) {
	l := new(LoginDB)
	err := r.dbHandler.Db.Where("username = ?", l.Username).First(&l).Error
	if err != nil {
		return false, err
	}

	return new(LoginDB) != l, nil
}

func (r *LoginRepository) Get(u username.Username) (loginModel.Login, error) {
	login := new(LoginDB)
	err := r.dbHandler.Db.Where("username = ?", login.Username).First(&login).Error
	if err != nil {
		return loginModel.Login{}, err
	}

	l, err := toLogin(*login)
	if err.Error() == username.InvalidUsername {
		return loginModel.Login{}, fmt.Errorf("user not found")
	}

	return l, nil
}
