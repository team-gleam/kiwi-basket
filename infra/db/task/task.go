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
	Username string
	Date time.Time
	Title string
}

func transformTaskForDB(t taskModel.Task, u username.Username) taskDB {
	return taskDB{uint(t.ID()), u.Name(), t.Date(), t.Title()}
}

func toTask(t taskDB) (taskModel.Task, username.Username, error) {
	task, err := taskModel.NewTask(int(t.ID), t.Date.Format(taskModel.Layout), t.Title)
	if err != nil {
		return taskModel.Task{}, username.Username{}, err
	}

	u, err := username.NewUsername(t.Username)
	return task, u, err
}

func (r *TaskRepository) Create(u username.Username, t taskModel.Task) error {
	d := transformTaskForDB(t, u)
	return r.dbHandler.Db.Create(d).Error
}

func (r *TaskRepository) GetAll(u username.Username) ([]taskModel.Task, error) {
	ds := make([]taskDB, 0)
	err := r.dbHandler.Db.Where("username = ?", u.Name).Find(&ds).Error
	if err != nil {
		return []taskModel.Task{}, err
	}

	tasks := make([]taskModel.Task, 0)
	for _, d := range ds {
		t, _, err := toTask(d)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (r *TaskRepository) Remove(u username.Username, id int) error {
	return r.dbHandler.Db.Where("username = ?", u.Name).Delete(taskDB{ID : uint(id)}).Error
}