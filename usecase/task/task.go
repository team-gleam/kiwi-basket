package task

import (
	"fmt"

	taskModel "github.com/team-gleam/kiwi-basket/domain/model/task"
	tokenModel "github.com/team-gleam/kiwi-basket/domain/model/user/token"
	taskRepository "github.com/team-gleam/kiwi-basket/domain/repository/task"
	credentialUsecase "github.com/team-gleam/kiwi-basket/usecase/user/credential"
)

type TaskUsecase struct {
	credentialUsecase credentialUsecase.CredentialUsecase
	taskRepository    taskRepository.ITaskRepository
}

func NewTaskUsecase(c credentialUsecase.CredentialUsecase, t taskRepository.ITaskRepository) TaskUsecase {
	return TaskUsecase{c, t}
}

const (
	IDIsNotZero = "ID is not zero"
	InvalidID   = "Invalid ID"
)

func (u TaskUsecase) Add(token tokenModel.Token, task taskModel.Task) error {
	credentialed, err := u.credentialUsecase.IsCredentialed(token)
	if err != nil {
		return err
	}
	if !credentialed {
		return fmt.Errorf(credentialUsecase.InvalidToken)
	}

	user, err := u.credentialUsecase.Whose(token)
	if err != nil {
		return err
	}

	return u.taskRepository.Create(user, task)
}

func (u TaskUsecase) Delete(token tokenModel.Token, id int) error {
	credentialed, err := u.credentialUsecase.IsCredentialed(token)
	if err != nil {
		return err
	}
	if !credentialed {
		return fmt.Errorf(credentialUsecase.InvalidToken)
	}

	if id == 0 {
		return fmt.Errorf(IDIsNotZero)
	}
	if id < 0 {
		return fmt.Errorf(InvalidID)
	}

	user, err := u.credentialUsecase.Whose(token)
	if err != nil {
		return err
	}

	return u.taskRepository.Remove(user, id)
}

func (u TaskUsecase) GetAll(token tokenModel.Token) ([]taskModel.Task, error) {
	credentialed, err := u.credentialUsecase.IsCredentialed(token)
	if err != nil {
		return nil, err
	}
	if !credentialed {
		return nil, fmt.Errorf(credentialUsecase.InvalidToken)
	}

	user, err := u.credentialUsecase.Whose(token)
	if err != nil {
		return nil, err
	}

	return u.taskRepository.GetAll(user)
}
