package credential

import (
	"fmt"

	credentialModel "github.com/team-gleam/kiwi-basket/domain/model/user/credential"
	loginModel "github.com/team-gleam/kiwi-basket/domain/model/user/login"
	"github.com/team-gleam/kiwi-basket/domain/model/user/token"
	"github.com/team-gleam/kiwi-basket/domain/model/user/username"
	credentialRepository "github.com/team-gleam/kiwi-basket/domain/repository/user/credential"
	loginRepository "github.com/team-gleam/kiwi-basket/domain/repository/user/login"
	loginUsecase "github.com/team-gleam/kiwi-basket/usecase/user/login"
)

type CredentialUsecase struct {
	credentialRepository credentialRepository.ICredentialRepository
	loginUsecase         loginUsecase.LoginUsecase
}

func NewCredentialUsecase(
	c credentialRepository.ICredentialRepository,
	l loginRepository.ILoginRepository,
) CredentialUsecase {
	return CredentialUsecase{
		c,
		loginUsecase.NewLoginUsecase(l),
	}
}

const (
	UserNotFound              = "user not found"
	InvalidUsernameOrPassword = "invalid username or password"
	InvalidToken              = "invalid token"
)

func (u CredentialUsecase) Generate(login loginModel.Login) (token.Token, error) {
	verified, err := u.loginUsecase.Verify(login)
	if err != nil {
		return token.NewToken(""), err
	}
	if !verified {
		return token.NewToken(""), fmt.Errorf(InvalidUsernameOrPassword)
	}

	t, err := token.GenToken()
	if err != nil {
		return token.NewToken(""), err
	}

	a := credentialModel.NewAuth(login.Username(), t)

	err = u.credentialRepository.Remove(a.Username())
	if err != nil {
		return token.NewToken(""), err
	}

	return t, u.credentialRepository.Append(a)
}

func (u CredentialUsecase) Delete(login loginModel.Login) error {
	verified, err := u.loginUsecase.Verify(login)
	if err != nil {
		return err
	}
	if !verified {
		return fmt.Errorf(InvalidUsernameOrPassword)
	}

	return u.credentialRepository.Remove(login.Username())
}

func (u CredentialUsecase) Get(login loginModel.Login) (credentialModel.Auth, error) {
	verified, err := u.loginUsecase.Verify(login)
	if err != nil {
		return credentialModel.Auth{}, err
	}
	if !verified {
		return credentialModel.Auth{}, fmt.Errorf(InvalidUsernameOrPassword)
	}

	return u.credentialRepository.GetByUsername(login.Username())
}

func (u CredentialUsecase) HasCredential(t token.Token) (bool, error) {
	return u.credentialRepository.Exists(t)
}

func (u CredentialUsecase) Whose(t token.Token) (username.Username, error) {
	a, err := u.credentialRepository.GetByToken(t)
	return a.Username(), err
}
