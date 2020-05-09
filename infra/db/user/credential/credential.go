package credential

import (
	"fmt"

	"github.com/jinzhu/gorm"
	credentialModel "github.com/team-gleam/kiwi-basket/domain/model/user/credential"
	"github.com/team-gleam/kiwi-basket/domain/model/user/token"
	"github.com/team-gleam/kiwi-basket/domain/model/user/username"
	credentialRepository "github.com/team-gleam/kiwi-basket/domain/repository/user/credential"
	"github.com/team-gleam/kiwi-basket/infra/db/handler"
)

type CredentialRepository struct {
	dbHandler *handler.DbHandler
}

func NewCredentialRepository(h *handler.DbHandler) credentialRepository.ICredentialRepository {
	h.Db.AutoMigrate(AuthDB{})
	return &CredentialRepository{h}
}

type AuthDB struct {
	Username string `gorm:"primary_key"`
	Token    string `gorm:"primary_key"`
}

func transformAuthForDB(a credentialModel.Auth) AuthDB {
	return AuthDB{a.Username().Name(), a.Token().Token()}
}

func toAuth(a AuthDB) (credentialModel.Auth, error) {
	u, err := username.NewUsername(a.Username)
	return credentialModel.NewAuth(u, token.NewToken(a.Token)), err
}

func (r *CredentialRepository) Append(a credentialModel.Auth) error {
	d := transformAuthForDB(a)
	return r.dbHandler.Db.Create(&d).Error
}

func (r *CredentialRepository) Remove(a credentialModel.Auth) error {
	d := transformAuthForDB(a)
	err := r.dbHandler.Db.Where("username = ?", d.Username).Delete(AuthDB{}).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	fmt.Println(err)
	return err
}

func (r *CredentialRepository) Exists(t token.Token) (bool, error) {
	a := new(AuthDB)
	err := r.dbHandler.Db.Where("token = ?", t.Token()).Take(a).Error
	if gorm.IsRecordNotFoundError(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return a.Token != "", nil
}

func (r *CredentialRepository) Get(t token.Token) (credentialModel.Auth, error) {
	auth := new(AuthDB)
	err := r.dbHandler.Db.Where("token = ?", t.Token()).Take(auth).Error
	if err != nil {
		return credentialModel.Auth{}, err
	}

	a, err := toAuth(*auth)
	if err != nil && err.Error() == username.InvalidUsername {
		return credentialModel.Auth{}, fmt.Errorf("user not found")
	}

	return a, nil
}
