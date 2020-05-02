package timetables

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
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
	Timetables TimetablesJSON `json:"timetable" validate:"required"`
}

type TimetablesJSON struct {
	Mon TimetableJSON `json:"mon" validate:"required"`
	Tue TimetableJSON `json:"tue" validate:"required"`
	Wed TimetableJSON `json:"wed" validate:"required"`
	Thu TimetableJSON `json:"thu" validate:"required"`
	Fri TimetableJSON `json:"fri" validate:"required"`
}

type TimetableJSON struct {
	One   *ClassJSON `json:"1"`
	Two   *ClassJSON `json:"2"`
	Three *ClassJSON `json:"3"`
	Four  *ClassJSON `json:"4"`
	Five  *ClassJSON `json:"5"`
}

type ClassJSON struct {
	Subject string  `json:"subject" validate:"max=85"`
	Room    *string `json:"room" validate:"omitempty,max_85_ptr|isdefault"`
}

func (t TimetablesResponse) IsValidated() (bool, error) {
	v := validator.New()
	err := v.RegisterValidation("max_85_ptr", Max85Ptr)
	if err != nil {
		return false, err
	}

	return v.Struct(t) == nil, nil
}

func Max85Ptr(validate validator.FieldLevel) bool {
	return len(validate.Field().String()) < 86
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

	if t.Room == nil {
		return timetablesModel.NoRoom(t.Subject)
	}

	return timetablesModel.NewClass(t.Subject, *t.Room)
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

	if c.IsNoRoom() {
		return &ClassJSON{
			Subject: c.Subject(),
			Room:    nil,
		}
	}

	room := c.Room()
	return &ClassJSON{
		Subject: c.Subject(),
		Room:    &room,
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
			errorResponse.NewError(fmt.Errorf(loginController.InvalidJSONFormat)),
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
