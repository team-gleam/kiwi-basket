package login

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	loginModel "github.com/team-gleam/kiwi-basket/domain/model/user/login"
	"github.com/team-gleam/kiwi-basket/domain/model/user/username"
	loginRepository "github.com/team-gleam/kiwi-basket/domain/repository/user/login"
	errorResponse "github.com/team-gleam/kiwi-basket/interfaces/controllers/error"
	loginUsecase "github.com/team-gleam/kiwi-basket/usecase/user/login"
)

type LoginController struct {
	loginUsecase loginUsecase.LoginUsecase
}

func NewLoginController(r loginRepository.ILoginRepository) *LoginController {
	return &LoginController{
		loginUsecase.NewLoginUsecase(r),
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

	err = c.loginUsecase.Delete(l)
	if err != nil && err.Error() == loginUsecase.UsernameNotFound {
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

	return ctx.NoContent(http.StatusOK)
}
