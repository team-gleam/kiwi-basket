package timetables

import (
	timetablesModel "github.com/team-gleam/kiwi-basket/domain/model/timetables"
	timetablesRepository "github.com/team-gleam/kiwi-basket/domain/repository/timetables"
	credentialRepository "github.com/team-gleam/kiwi-basket/domain/repository/user/credential"
	loginRepository "github.com/team-gleam/kiwi-basket/domain/repository/user/login"
	timetablesUsecase "github.com/team-gleam/kiwi-basket/usecase/timetables"
	credentialUsecase "github.com/team-gleam/kiwi-basket/usecase/user/credential"
)

type TimetablesController struct {
	timetablesUsecase timetablesUsecase.TimetablesUsecase
}

func NewTimetablesController(
	c credentialRepository.ICredentialRepository,
	l loginRepository.ILoginRepository,
	t timetablesRepository.ITimetablesRepository,
) *TimetablesController {
	return &TimetablesController{
		timetablesUsecase.NewTimetablesUsecase(
			credentialUsecase.NewCredentialUsecase(c, l),
			t,
		),
	}
}

type TimetablesResponse struct {
	Timetables Timetables `json:"timetable"`
}

type Timetables struct {
	Mon Timetable `json:"mon"`
	Tue Timetable `json:"tue"`
	Wed Timetable `json:"wed"`
	Thu Timetable `json:"thu"`
	Fri Timetable `json:"fri"`
}

type Timetable struct {
	One   *Class `json:"1"`
	Two   *Class `json:"2"`
	Three *Class `json:"3"`
	Four  *Class `json:"4"`
	Five  *Class `json:"5"`
}

type Class struct {
	Subject string `json:"subject"`
	Room    string `json:"room"`
}

func (t TimetablesResponse) toTimetables() timetablesModel.Timetables {
	return timetablesModel.NewTimetables(
		t.Timetables.Mon.toTimetable(),
		t.Timetables.Tue.toTimetable(),
		t.Timetables.Wed.toTimetable(),
		t.Timetables.Thu.toTimetable(),
		t.Timetables.Fri.toTimetable(),
	)
}

func (t Timetable) toTimetable() timetablesModel.Timetable {
	return timetablesModel.NewTimetable(
		t.One.toClass(),
		t.Two.toClass(),
		t.Three.toClass(),
		t.Four.toClass(),
		t.Five.toClass(),
	)
}

func (t *Class) toClass() timetablesModel.Class {
	if t == nil {
		return timetablesModel.NoClass()
	}
	return timetablesModel.NewClass(t.Subject, t.Room)
}
