package task

import (
	"time"

	taskModel "github.com/team-gleam/kiwi-basket/domain/model/task"
)

type taskDB struct {
	ID uint `gorm:"primary_key;auto_increment"`
	username string
	date time.Time
	title string
}

func transformTaskForDB(t taskModel.Task, u string) taskDB {
	return taskDB{uint(t.ID()), u, t.Date(), t.Title()}
}

func toTask(t taskDB) (taskModel.Task, string, error) {
	task, err := taskModel.NewTask(int(t.ID), t.date.Format(taskModel.Layout), t.title)
	return task, t.username, err
}