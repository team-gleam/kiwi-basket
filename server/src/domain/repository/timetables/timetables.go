package timetables

import (
	"github.com/team-gleam/kiwi-basket/server/src/domain/model/timetables"
	"github.com/team-gleam/kiwi-basket/server/src/domain/model/user/username"
)

type ITimetablesRepository interface {
	Create(username.Username, timetables.Timetables) error
	Delete(username.Username) error
	Exists(username.Username) (bool, error)
	Get(username.Username) (timetables.Timetables, error)
}
