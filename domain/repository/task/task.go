package task

import (
	"github.com/the-gleam/kiwi-basket/domain/model/task"
	"github.com/the-gleam/kiwi-basket/domain/model/user/username"
)

type ITaskRepository interface {
	Create(username.Username, task.Task) error
	GetAll(username.Username) ([]task.Task, error)
}
