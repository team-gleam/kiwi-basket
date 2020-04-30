package credential

import (
	"fmt"

	credentialModel "github.com/team-gleam/kiwi-basket/domain/model/user/credential"
	loginModel "github.com/team-gleam/kiwi-basket/domain/model/user/login"
	"github.com/team-gleam/kiwi-basket/domain/model/user/token"
	"github.com/team-gleam/kiwi-basket/domain/model/user/username"
	credentialRepository "github.com/team-gleam/kiwi-basket/domain/repository/user/credential"
	loginRepository "github.com/team-gleam/kiwi-basket/domain/repository/user/login"
)

type CredentialUsecase struct {
	credentialRepository credentialRepository.ICredentialRepository
	loginRepository      loginRepository.ILoginRepository
}

func NewCredentialUsecase(
	c credentialRepository.ICredentialRepository,
	l loginRepository.ILoginRepository,
) CredentialUsecase {
	return CredentialUsecase{c, l}
}

const (
	UserNotFound              = "user not found"
	InvalidUsernameOrPassword = "invalid username or password"
	InvalidToken              = "invalid token"
)

func (u CredentialUsecase) Generate(login loginModel.Login) (token.Token, error) {
	exist, err := u.loginRepository.Exists(login.Username())
	if err != nil {
		return token.NewToken(""), err
	}
	if !exist {
		return token.NewToken(""), fmt.Errorf(UserNotFound)
	}

	l, err := u.loginRepository.Get(login.Username())
	if err != nil {
		return token.NewToken(""), err
	}
	if l != login {
		return token.NewToken(""), fmt.Errorf(InvalidUsernameOrPassword)
	}

	t, err := token.GenToken()
	if err != nil {
		return token.NewToken(""), err
	}

	a := credentialModel.NewAuth(login.Username(), t)
	return t, u.credentialRepository.Append(a)
}

func (u CredentialUsecase) IsCredentialed(t token.Token) (bool, error) {
	return u.credentialRepository.Exists(t)
}

func (u CredentialUsecase) Whose(t token.Token) (username.Username, error) {
	a, err := u.credentialRepository.Get(t)
	return a.Username(), err
}
