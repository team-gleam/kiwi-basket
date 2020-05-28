package login

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	loginModel "github.com/team-gleam/kiwi-basket/server/src/domain/model/user/login"
	"github.com/team-gleam/kiwi-basket/server/src/domain/model/user/username"
	taskRepository "github.com/team-gleam/kiwi-basket/server/src/domain/repository/task"
	timetablesRepository "github.com/team-gleam/kiwi-basket/server/src/domain/repository/timetables"
	credentialRepository "github.com/team-gleam/kiwi-basket/server/src/domain/repository/user/credential"
	loginRepository "github.com/team-gleam/kiwi-basket/server/src/domain/repository/user/login"
	errorResponse "github.com/team-gleam/kiwi-basket/server/src/interfaces/controllers/error"
	taskUsecase "github.com/team-gleam/kiwi-basket/server/src/usecase/task"
	timetablesUsecase "github.com/team-gleam/kiwi-basket/server/src/usecase/timetables"
	credentialUsecase "github.com/team-gleam/kiwi-basket/server/src/usecase/user/credential"
	loginUsecase "github.com/team-gleam/kiwi-basket/server/src/usecase/user/login"
)

type LoginController struct {
	loginUsecase      loginUsecase.LoginUsecase
	credentialUsecase credentialUsecase.CredentialUsecase
	taskUsecase       taskUsecase.TaskUsecase
	timetablesUsecase timetablesUsecase.TimetablesUsecase
}

func NewLoginController(
	l loginRepository.ILoginRepository,
	c credentialRepository.ICredentialRepository,
	t taskRepository.ITaskRepository,
	tt timetablesRepository.ITimetablesRepository,
) *LoginController {
	return &LoginController{
		loginUsecase.NewLoginUsecase(l),
		credentialUsecase.NewCredentialUsecase(c, l),
		taskUsecase.NewTaskUsecase(c, l, t),
		timetablesUsecase.NewTimetablesUsecase(c, l, tt),
	}
}

const (
	InvalidUsernameOrPassword = "invalid username or password"
	InvalidJSONFormat         = "invalid JSON format"
)

type LoginResponse struct {
	Username string `json:"username" validate:"required,alphanum,max=255"`
	Password string `json:"password" validate:"required,alphanum,min=8,max=72"`
}

func (l LoginResponse) Validates() bool {
	return validator.New().Struct(l) == nil
}

func (l LoginResponse) ToLogin() (loginModel.Login, error) {
	u, err := username.NewUsername(l.Username)
	if err != nil {
		return loginModel.Login{}, err
	}

	hashed := hashPassword(l.Password)

	return loginModel.NewLogin(u, hashed), nil
}

func hashPassword(p string) string {
	b := sha256.Sum256([]byte(p))
	return hex.EncodeToString(b[:])
}

func (c LoginController) SignUp(ctx echo.Context) error {
	login := new(LoginResponse)
	err := ctx.Bind(login)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			errorResponse.NewError(fmt.Errorf(InvalidJSONFormat)),
		)
	}

	if !login.Validates() {
		return ctx.JSON(
			http.StatusBadRequest,
			errorResponse.NewError(fmt.Errorf(InvalidUsernameOrPassword)),
		)
	}

	l, err := login.ToLogin()
	if err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			errorResponse.NewError(fmt.Errorf(errorResponse.InternalServerError)),
		)
	}

	err = c.loginUsecase.Add(l)
	if err != nil && err.Error() == loginUsecase.UsernameAlreadyExists {
		return ctx.JSON(
			http.StatusConflict,
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

func (c LoginController) DeleteAccound(ctx echo.Context) error {
	login := new(LoginResponse)
	err := ctx.Bind(login)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			errorResponse.NewError(fmt.Errorf(InvalidJSONFormat)),
		)
	}

	if !login.Validates() {
		return ctx.JSON(
			http.StatusBadRequest,
			errorResponse.NewError(fmt.Errorf(InvalidJSONFormat)),
		)
	}

	l, err := login.ToLogin()
	if err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			errorResponse.NewError(fmt.Errorf(errorResponse.InternalServerError)),
		)
	}

	verified, err := c.loginUsecase.Verify(l)
	if err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			errorResponse.NewError(fmt.Errorf(errorResponse.InternalServerError)),
		)
	}
	if !verified {
		return ctx.JSON(
			http.StatusUnauthorized,
			errorResponse.NewError(fmt.Errorf(InvalidUsernameOrPassword)),
		)
	}

	auth, err := c.credentialUsecase.Get(l)
	if err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			errorResponse.NewError(fmt.Errorf(errorResponse.InternalServerError)),
		)
	}

	token := auth.Token()

	if err = c.taskUsecase.DeleteAll(token); err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			errorResponse.NewError(fmt.Errorf(errorResponse.InternalServerError)),
		)
	}

	if err = c.timetablesUsecase.Delete(token); err != nil && err.Error() != timetablesUsecase.TimetablesNotFound {
		return ctx.JSON(
			http.StatusInternalServerError,
			errorResponse.NewError(fmt.Errorf(errorResponse.InternalServerError)),
		)
	}

	if err = c.credentialUsecase.Delete(l); err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			errorResponse.NewError(fmt.Errorf(errorResponse.InternalServerError)),
		)
	}

	err = c.loginUsecase.Delete(l)
	if err != nil && err.Error() == loginUsecase.UsernameNotFound {
		return ctx.JSON(
			http.StatusUnauthorized,
			errorResponse.NewError(fmt.Errorf(InvalidUsernameOrPassword)),
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
