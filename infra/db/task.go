package task

import (
	"time"

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