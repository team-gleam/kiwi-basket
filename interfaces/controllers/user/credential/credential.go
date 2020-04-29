package credential

import (
	credentialRepository "github.com/team-gleam/kiwi-basket/model/repository/user/credential"
	loginRepository "github.com/team-gleam/kiwi-basket/model/repository/user/login"
	credentialUsecase "github.com/team-gleam/kiwi-basket/usecase/user/credential"
)

type CredentialController struct {
	credentialUsecase credentialUsecase.CredentialUsecase
}

func NewCredentialController(
	c credentialRepository.ICredentialRepository,
	l loginRepository.ILoginRepository,
) *CredentialController {
	return &CredentialController{
		credentialUsecase.NewCredentialRepository(c, l),
	}
}
