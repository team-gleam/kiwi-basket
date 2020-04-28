package credential

import (
	"fmt"

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
	return &CredentialRepository{h}
}

type authDB struct {
	Username string `gorm:"primary_key"`
	Token    string `gorm:"primary_key"`
}

func transformAuthForDB(a credentialModel.Auth) authDB {
	return authDB{a.Username().Name(), a.Token().Token()}
}

func toAuth(a authDB) (credentialModel.Auth, error) {
	u, err := username.NewUsername(a.Username)
	return credentialModel.NewAuth(u, token.NewToken(a.Token)), err
}

func (r *CredentialRepository) Append(a credentialModel.Auth) error {
	d := transformAuthForDB(a)
	return r.dbHandler.Db.Create(d).Error
}

func (r *CredentialRepository) Remove(a credentialModel.Auth) error {
	d := transformAuthForDB(a)
	return r.dbHandler.Db.Delete(d).Error
}

func (r *CredentialRepository) Exists(t token.Token) (bool, error) {
	a := new(authDB)
	err := r.dbHandler.Db.Where("token = ?", t.Token).First(&a).Error
	if err != nil {
		return false, err
	}
	return new(authDB) != a, nil
}

func (r *CredentialRepository) Get(t token.Token) (credentialModel.Auth, error) {
	auth := new(authDB)
	err := r.dbHandler.Db.Where("token = ?", t.Token).First(&auth).Error
	if err != nil {
		return credentialModel.Auth{}, err
	}

	a, err := toAuth(*auth)
	if err.Error() == username.InvalidUsername {
		return credentialModel.Auth{}, fmt.Errorf("user not found")
	}

	return a, nil
}
