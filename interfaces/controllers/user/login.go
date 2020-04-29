package login

import (
	loginRepository "github.com/team-gleam/kiwi-basket/domain/repository/user/login"
	loginUsecase "github.com/team-gleam/kiwi-basket/usecase/user/login"
)

type LoginController struct {
	loginUsecase loginUsecase.LoginUsecase
}

func NewLoginController(r loginRepository.ILoginRepository) *LoginController {
	return &LoginController{
		loginUsecase.NewLoginUsecase(r),
	}
}
