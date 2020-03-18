package credential

import (
	"fmt"

	"github.com/the-gleam/kiwi-basket/domain/model/user/credential"
	loginModel "github.com/the-gleam/kiwi-basket/domain/model/user/login"
	"github.com/the-gleam/kiwi-basket/domain/model/user/token"
	credentialRepository "github.com/the-gleam/kiwi-basket/domain/repository/user/credential"
	"github.com/the-gleam/kiwi-basket/domain/repository/user/login"
	loginRepository "github.com/the-gleam/kiwi-basket/domain/repository/user/login"
)

type CredentialUsecase struct {
	credentialRepository credentialRepository.ICredentialRepository
	loginRepository      loginRepository.ILoginRepository
}

func NewCredentialRepository(c credentialRepository.ICredentialRepository, l login.ILoginRepository) CredentialUsecase {
	return CredentialUsecase{c, l}
}

func (u CredentialUsecase) Generate(login loginModel.Login) (token.Token, error) {
	exist, err := u.loginRepository.Exists(login.Username())
	if err != nil {
		return token.NewToken(""), err
	}
	if !exist {
		return token.NewToken(""), fmt.Errorf("this user not exists")
	}

	l, err := u.loginRepository.Get(login.Username())
	if err != nil {
		return token.NewToken(""), err
	}
	if l.HashedPassword() != login.HashedPassword() {
		return token.NewToken(""), fmt.Errorf("username or password is invalid")
	}

	t, err := token.GenToken()
	if err != nil {
		return token.NewToken(""), err
	}

	a := credential.NewAuth(login.Username(), t)
	return t, u.credentialRepository.Append(a)
}

func (u CredentialUsecase) Delete(t token.Token) error {
	exist, err := u.credentialRepository.Exists(t)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("this token not exists")
	}

	a, err := u.credentialRepository.Get(t)
	if err != nil {
		return err
	}

	return u.credentialRepository.Remove(a)
}
