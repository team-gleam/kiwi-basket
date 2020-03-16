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
	exist, err := u.loginRepository.Exists(l)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("this username already exists")
	}

	err = u.loginRepository.Create(l)
	if err != nil {
		return err
	}

	return nil
}
