package timetables

import (
	"github.com/team-gleam/kiwi-basket/domain/model/timetables"
	"github.com/team-gleam/kiwi-basket/domain/model/user/username"
)

type ITimetablesRepository interface {
	Create(username.Username, timetables.Timetables) error
	Delete(username.Username) error
	Exist(username.Username) (bool, error)
	Get(username.Username) (timetables.Timetables, error)
}
