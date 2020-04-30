package timetables

import (
	timetablesRepository "github.com/team-gleam/kiwi-basket/domain/repository/timetables"
	credentialRepository "github.com/team-gleam/kiwi-basket/domain/repository/user/credential"
	loginRepository "github.com/team-gleam/kiwi-basket/domain/repository/user/login"
	timetablesUsecase "github.com/team-gleam/kiwi-basket/usecase/timetables"
	credentialUsecase "github.com/team-gleam/kiwi-basket/usecase/user/credential"
)

type TimetablesController struct {
	timetablesUsecase timetablesUsecase.TimetablesUsecase
}

func NewTimetablesController(
	c credentialRepository.ICredentialRepository,
	l loginRepository.ILoginRepository,
	t timetablesRepository.ITimetablesRepository,
) *TimetablesController {
	return &TimetablesController{
		timetablesUsecase.NewTimetablesUsecase(
			credentialUsecase.NewCredentialUsecase(c, l),
			t,
		),
	}
}
