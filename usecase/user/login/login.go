package login

import (
	"fmt"

	loginModel "github.com/the-gleam/kiwi-basket/domain/model/user/login"
	loginRepository "github.com/the-gleam/kiwi-basket/domain/repository/user/login"
)

type LoginUsecase struct {
	loginRepository loginRepository.ILoginRepository
}

func NewLoginUsecase(r loginRepository.ILoginRepository) LoginUsecase {
	return LoginUsecase{r}
}

func (u LoginUsecase) Add(l loginModel.Login) error {
	exist, err := u.loginRepository.Exists(l.Username())
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("this username already exists")
	}

	return u.loginRepository.Create(l)
}

func (u LoginUsecase) Delete(l loginModel.Login) error {
	exist, err := u.loginRepository.Exists(l.Username())
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("this username not exists")
	}

	return u.loginRepository.Delete(l)
}
