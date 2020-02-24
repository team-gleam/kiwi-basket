package task

import (
	"fmt"
	"time"
)

type Task struct {
	date  time.Time
	title string
}

const (
	layout = "2006-01-02"
)

func NewTask(date, title string) (Task, error) {

	d, err := time.Parse(layout, date)
	if err != nil {
		return Task{}, fmt.Errorf("invalid date format")
	}

	return Task{d, title}, nil
}

func (t Task) TextDate() string {
	return t.date.Format(layout)
}

func (t Task) Title() string {
	return t.title
}
