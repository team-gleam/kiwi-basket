package timetables

import (
	"github.com/the-gleam/kiwi-basket/domain/model/timetables"
	"github.com/the-gleam/kiwi-basket/domain/model/user/username"
)

type ITimetablesRepository interface {
	Create(username.Username, timetables.Timetables) error
	Delete(username.Username) error
	Get(username.Username) (timetables.Timetables, error)
}
