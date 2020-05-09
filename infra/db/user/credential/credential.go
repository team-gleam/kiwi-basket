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
	h.Db.AutoMigrate(CredentialDB{})
	return &CredentialRepository{h}
}

type CredentialDB struct {
	Username string `gorm:"primary_key"`
	Token    string `gorm:"primary_key"`
}

func transformAuthForDB(a credentialModel.Auth) CredentialDB {
	return CredentialDB{a.Username().Name(), a.Token().Token()}
}

func toAuth(a CredentialDB) (credentialModel.Auth, error) {
	u, err := username.NewUsername(a.Username)
	return credentialModel.NewAuth(u, token.NewToken(a.Token)), err
}

func (r *CredentialRepository) Append(a credentialModel.Auth) error {
	d := transformAuthForDB(a)
	return r.dbHandler.Db.Create(&d).Error
}

func (r *CredentialRepository) Remove(a credentialModel.Auth) error {
	d := transformAuthForDB(a)
	return r.dbHandler.Db.Delete(&d).Error
}

func (r *CredentialRepository) Exists(t token.Token) (bool, error) {
	a := new(CredentialDB)
	err := r.dbHandler.Db.Where("token = ?", t.Token).Take(&a).Error
	if gorm.IsRecordNotFoundError(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return new(CredentialDB) != a, nil
}

func (r *CredentialRepository) Get(t token.Token) (credentialModel.Auth, error) {
	auth := new(CredentialDB)
	err := r.dbHandler.Db.Where("token = ?", t.Token).Take(&auth).Error
	if err != nil {
		return credentialModel.Auth{}, err
	}

	a, err := toAuth(*auth)
	if err.Error() == username.InvalidUsername {
		return credentialModel.Auth{}, fmt.Errorf("user not found")
	}

	return a, nil
}
