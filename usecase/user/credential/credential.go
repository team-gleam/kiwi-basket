package credential

import (
	"fmt"

	"github.com/the-gleam/kiwi-basket/domain/model/user/credential"
	loginModel "github.com/the-gleam/kiwi-basket/domain/model/user/login"
	"github.com/the-gleam/kiwi-basket/domain/model/user/token"
	"github.com/the-gleam/kiwi-basket/domain/model/user/username"
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
	if l != login {
		return token.NewToken(""), fmt.Errorf("username or password is invalid")
	}

	t, err := token.GenToken()
	if err != nil {
		return token.NewToken(""), err
	}

	a := credential.NewAuth(login.Username(), t)
	return t, u.credentialRepository.Append(a)
}

func (u CredentialUsecase) IsCredentialed(t token.Token) (bool, error) {
	return u.credentialRepository.Exists(t)
}

func (u CredentialUsecase) Whose(t token.Token) (username.Username, error) {
	a, err := u.credentialRepository.Get(t)
	if err != nil {
		return username.Username{}, err
	}
	return a.Username(), nil
}
