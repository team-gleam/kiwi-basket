package task

import (
	"fmt"
	"time"
)

type Task struct {
	id    int
	date  time.Time
	title string
}

const (
	Layout = "2006-01-02"
)

func NewTask(id int, date, title string) (Task, error) {

	d, err := time.Parse(Layout, date)
	if err != nil {
		return Task{}, fmt.Errorf("invalid date format")
	}

	return Task{id, d, title}, nil
}

func (t Task) ID() int {
	return t.id
}

func (t Task) TextDate() string {
	return t.date.Format(Layout)
}

func (t Task) Date() time.Time {
	return t.date
}

func (t Task) Title() string {
	return t.title
}
