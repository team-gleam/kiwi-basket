package task

import (
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {
	tests := []struct {
		name       string
		dateString string
		shouldFail bool
		expected   time.Time
	}{
		{"splited by -", "2020-01-01", false, time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)},
		{"0 deleted", "2020-1-1", true, time.Now()},
		{"splited by /", "2020/01/01", true, time.Now()},
		{"include time", "2020/01/01 0:00", true, time.Now()},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v, e := NewTask(0, test.dateString, "")

			if test.shouldFail && e != nil {
				return
			} else if !test.shouldFail && e != nil {
				t.Fatalf("unexpected error: %v", e)
			} else if test.shouldFail && e == nil {
				t.Fatalf("expected error but got nil")
			} else if test.expected != v.date {
				t.Fatalf("expected: %v; got: %v\n", test.expected, v)
			}
		})
	}
}

func TestTextDate(t *testing.T) {
	tests := []struct {
		name string
		task Task
	}{
		{"2020-01-01", Task{0, time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC), ""}},
		{"2020-11-01", Task{0, time.Date(2020, time.November, 1, 0, 0, 0, 0, time.UTC), ""}},
		{"2020-01-11", Task{0, time.Date(2020, time.January, 11, 0, 0, 0, 0, time.UTC), ""}},
		{"2020-11-11", Task{0, time.Date(2020, time.November, 11, 0, 0, 0, 0, time.UTC), ""}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := test.task.TextDate()

			if test.name != v {
				t.Fatalf("expected: %v; got: %v\n", test.name, v)
			}
		})
	}
}

func TestTaskGetters(t *testing.T) {
	task, _ := NewTask(1, "2020-01-01", "title")
	tests := []struct {
		expected interface{}
		got      interface{}
	}{
		{task.id, task.ID()},
		{task.date, task.Date()},
		{task.title, task.Title()},
	}

	for _, test := range tests {
		if test.expected != test.got {
			t.Fatalf(
				"expected: %v; got: %v\n",
				test.expected,
				test.got,
			)
		}
	}
}
