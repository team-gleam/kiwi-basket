package task

import (
	taskRepository "github.com/team-gleam/kiwi-basket/domain/repository/task"
	credentialRepository "github.com/team-gleam/kiwi-basket/domain/repository/user/credential"
	loginRepository "github.com/team-gleam/kiwi-basket/domain/repository/user/login"
	taskUsecase "github.com/team-gleam/kiwi-basket/usecase/task"
	credentialUsecase "github.com/team-gleam/kiwi-basket/usecase/user/credential"
)

type TaskController struct {
	taskUsecase taskUsecase.TaskUsecase
}

func NewTaskController(
	c credentialRepository.ICredentialRepository,
	l loginRepository.ILoginRepository,
	t taskRepository.ITaskRepository,
) *TaskController {
	return &TaskController{
		taskUsecase.NewTaskUsecase(
			credentialUsecase.NewCredentialUsecase(c, l),
			t,
		),
	}
}
