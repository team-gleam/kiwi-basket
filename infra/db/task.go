package task

import (
	"time"

	taskModel "github.com/team-gleam/kiwi-basket/domain/model/task"
	"github.com/team-gleam/kiwi-basket/domain/model/user/username"
	taskRepository "github.com/team-gleam/kiwi-basket/domain/repository/task"
	"github.com/team-gleam/kiwi-basket/infra/db/handler"
)

type TaskRepository struct {
	dbHandler *handler.DbHandler
}

func NewTaskRepository(h *handler.DbHandler) taskRepository.ITaskRepository {
	return &TaskRepository{h}
}

type taskDB struct {
	ID uint `gorm:"primary_key;auto_increment"`
	username string
	date time.Time
	title string
}

func transformTaskForDB(t taskModel.Task, u username.Username) taskDB {
	return taskDB{uint(t.ID()), u.Name(), t.Date(), t.Title()}
}

func toTask(t taskDB) (taskModel.Task, username.Username, error) {
	task, err := taskModel.NewTask(int(t.ID), t.date.Format(taskModel.Layout), t.title)
	if err != nil {
		return taskModel.Task{}, username.Username{}, err
	}

	u, err := username.NewUsername(t.username)
	return task, u, err
}