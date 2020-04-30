package timetables

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	timetablesModel "github.com/team-gleam/kiwi-basket/domain/model/timetables"
	"github.com/team-gleam/kiwi-basket/domain/model/user/token"
	timetablesRepository "github.com/team-gleam/kiwi-basket/domain/repository/timetables"
	credentialRepository "github.com/team-gleam/kiwi-basket/domain/repository/user/credential"
	loginRepository "github.com/team-gleam/kiwi-basket/domain/repository/user/login"
	errorResponse "github.com/team-gleam/kiwi-basket/interfaces/controllers/error"
	loginController "github.com/team-gleam/kiwi-basket/interfaces/controllers/user/login"
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
	Timetables TimetablesJSON `json:"timetable"`
}

type TimetablesJSON struct {
	Mon TimetableJSON `json:"mon"`
	Tue TimetableJSON `json:"tue"`
	Wed TimetableJSON `json:"wed"`
	Thu TimetableJSON `json:"thu"`
	Fri TimetableJSON `json:"fri"`
}

type TimetableJSON struct {
	One   *ClassJSON `json:"1"`
	Two   *ClassJSON `json:"2"`
	Three *ClassJSON `json:"3"`
	Four  *ClassJSON `json:"4"`
	Five  *ClassJSON `json:"5"`
}

type ClassJSON struct {
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

func (t TimetableJSON) toTimetable() timetablesModel.Timetable {
	return timetablesModel.NewTimetable(
		t.One.toClass(),
		t.Two.toClass(),
		t.Three.toClass(),
		t.Four.toClass(),
		t.Five.toClass(),
	)
}

func (t *ClassJSON) toClass() timetablesModel.Class {
	if t == nil {
		return timetablesModel.NoClass()
	}
	return timetablesModel.NewClass(t.Subject, t.Room)
}

func toTimetablesResponse(t timetablesModel.Timetables) TimetablesResponse {
	return TimetablesResponse{
		Timetables: TimetablesJSON{
			Mon: toTimetableJSON(t.Mon()),
			Tue: toTimetableJSON(t.Tue()),
			Wed: toTimetableJSON(t.Wed()),
			Thu: toTimetableJSON(t.Thu()),
			Fri: toTimetableJSON(t.Fri()),
		},
	}
}

func toTimetableJSON(t timetablesModel.Timetable) TimetableJSON {
	return TimetableJSON{
		One:   toClassJSON(t.First()),
		Two:   toClassJSON(t.Second()),
		Three: toClassJSON(t.Third()),
		Four:  toClassJSON(t.Fourth()),
		Five:  toClassJSON(t.Fifth()),
	}
}

func toClassJSON(c timetablesModel.Class) *ClassJSON {
	if c.IsNoClass() {
		return nil
	}

	return &ClassJSON{
		Subject: c.Subject(),
		Room:    c.Room(),
	}
}

func (c TimetablesController) Register(ctx echo.Context) error {
	t := ctx.Request().Header.Get("Token")
	if t == "" {
		return ctx.JSON(
			http.StatusUnauthorized,
			errorResponse.NewError(fmt.Errorf(credentialUsecase.InvalidToken)),
		)
	}

	res := new(TimetablesResponse)
	err := ctx.Bind(res)
	if err != nil || res.Timetables.Mon.One == new(ClassJSON) {
		return ctx.JSON(
			http.StatusBadRequest,
			errorResponse.NewError(fmt.Errorf(loginController.InvalidJsonFormat)),
		)
	}

	timetables := res.toTimetables()

	err = c.timetablesUsecase.Add(token.NewToken(t), timetables)
	if err.Error() == credentialUsecase.InvalidToken {
		return ctx.JSON(
			http.StatusUnauthorized,
			errorResponse.NewError(err),
		)
	}
	if err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			errorResponse.NewError(fmt.Errorf(errorResponse.InternalServerError)),
		)
	}

	return ctx.NoContent(http.StatusOK)
}

func (c TimetablesController) Get(ctx echo.Context) error {
	t := ctx.Request().Header.Get("Token")
	if t == "" {
		return ctx.JSON(
			http.StatusUnauthorized,
			errorResponse.NewError(fmt.Errorf(credentialUsecase.InvalidToken)),
		)
	}

	timetables, err := c.timetablesUsecase.Get(token.NewToken(t))
	if err.Error() == credentialUsecase.InvalidToken {
		return ctx.JSON(
			http.StatusUnauthorized,
			errorResponse.NewError(err),
		)
	}
	if err.Error() == timetablesUsecase.TimetablesNotFound {
		return ctx.JSON(
			http.StatusNotFound,
			errorResponse.NewError(err),
		)
	}
	if err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			errorResponse.NewError(fmt.Errorf(errorResponse.InternalServerError)),
		)
	}

	res := toTimetablesResponse(timetables)

	return ctx.JSON(http.StatusOK, res)
}
