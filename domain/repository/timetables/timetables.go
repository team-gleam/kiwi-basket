package timetables

import (
	"github.com/the-gleam/kiwi-basket/domain/model/timetables"
)

type ITimetablesRepository interface {
	Create(timetables.Timetables) error
	Delete(timetables.Timetables) error
	Get() (timetables.Timetables, error)
}
