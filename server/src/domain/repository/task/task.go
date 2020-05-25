package task

import (
	"github.com/team-gleam/kiwi-basket/domain/model/task"
	"github.com/team-gleam/kiwi-basket/domain/model/user/username"
)

type ITaskRepository interface {
	Create(username.Username, task.Task) error
	GetAll(username.Username) ([]task.Task, error)
	Remove(username.Username, int) error
	RemoveAll(username.Username) error
}
