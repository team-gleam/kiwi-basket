package task

import "github.com/the-gleam/kiwi-basket/domain/model/task"

type ITaskRepository interface {
	Create(task.Task) error
	GetAll() ([]task.Task, error)
}
