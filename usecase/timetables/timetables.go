package timetables

import (
	"fmt"

	timetablesModel "github.com/team-gleam/kiwi-basket/domain/model/timetables"
	"github.com/team-gleam/kiwi-basket/domain/model/user/token"
	timetablesRepository "github.com/team-gleam/kiwi-basket/domain/repository/timetables"
	credentialUsecase "github.com/team-gleam/kiwi-basket/usecase/user/credential"
)

type TimetablesUsecase struct {
	credentialUsecase    credentialUsecase.CredentialUsecase
	timetablesRepository timetablesRepository.ITimetablesRepository
}

func NewTimetablesUsecase(c credentialUsecase.CredentialUsecase, t timetablesRepository.ITimetablesRepository) TimetablesUsecase {
	return TimetablesUsecase{c, t}
}

func (u TimetablesUsecase) Register(token token.Token, timetables timetablesModel.Timetables) error {
	credentialed, err := u.credentialUsecase.IsCredentialed(token)
	if err != nil {
		return err
	}
	if !credentialed {
		return fmt.Errorf("this token is not credentialed")
	}

	user, err := u.credentialUsecase.Whose(token)
	if err != nil {
		return err
	}

	exist, err := u.timetablesRepository.Exists(user)
	if err != nil {
		return err
	}
	if exist {
		if err = u.timetablesRepository.Delete(user); err != nil {
			return err
		}
	}

	return u.timetablesRepository.Create(user, timetables)
}

func (u TimetablesUsecase) Get(token token.Token) (timetablesModel.Timetables, error) {
	credentialed, err := u.credentialUsecase.IsCredentialed(token)
	if err != nil {
		return timetablesModel.Timetables{}, err
	}
	if !credentialed {
		return timetablesModel.Timetables{}, fmt.Errorf("this token is not credentialed")
	}

	user, err := u.credentialUsecase.Whose(token)
	if err != nil {
		return timetablesModel.Timetables{}, err
	}

	exist, err := u.timetablesRepository.Exists(user)
	if err != nil {
		return timetablesModel.Timetables{}, err
	}

	if !exist {
		return timetablesModel.Timetables{}, fmt.Errorf("timetables not found")
	}

	return u.timetablesRepository.Get(user)
}
