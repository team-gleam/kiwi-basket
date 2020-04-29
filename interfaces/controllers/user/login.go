package login

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	loginModel "github.com/team-gleam/kiwi-basket/domain/model/user/login"
	"github.com/team-gleam/kiwi-basket/domain/model/user/username"
	loginRepository "github.com/team-gleam/kiwi-basket/domain/repository/user/login"
	errorResponse "github.com/team-gleam/kiwi-basket/interfaces/controllers/error"
	loginUsecase "github.com/team-gleam/kiwi-basket/usecase/user/login"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	loginUsecase loginUsecase.LoginUsecase
}

func NewLoginController(r loginRepository.ILoginRepository) *LoginController {
	return &LoginController{
		loginUsecase.NewLoginUsecase(r),
	}
}

type LoginResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l LoginResponse) toLogin() (loginModel.Login, error) {
	u, err := username.NewUsername(l.Username)
	if err != nil {
		return loginModel.Login{}, err
	}

	hashed, err := hashPassword(l.Password)
	if err != nil {
		return loginModel.Login{}, err
	}

	return loginModel.NewLogin(u, hashed), nil
}

func hashPassword(p string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

const (
	InvalidUsernameOrPassword = "invalid username or password"
	InternalServerError       = "internal server error"
)

func (c LoginController) SignUp(ctx echo.Context) error {
	login := new(LoginResponse)
	err := ctx.Bind(login)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, errorResponse.NewError(err))
	}
	if login.Username == "" || login.Password == "" {
		return ctx.JSON(
			http.StatusBadRequest,
			errorResponse.NewError(fmt.Errorf(InvalidUsernameOrPassword)),
		)
	}

	l, err := login.toLogin()
	if err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			errorResponse.NewError(fmt.Errorf(InternalServerError)),
		)
	}

	err = c.loginUsecase.Add(l)
	if err.Error() == loginUsecase.UsernameAlreadyExists {
		return ctx.JSON(
			http.StatusConflict,
			errorResponse.NewError(fmt.Errorf(loginUsecase.UsernameAlreadyExists)),
		)
	}
	if err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			errorResponse.NewError(fmt.Errorf(InternalServerError)),
		)
	}

	return ctx.NoContent(http.StatusOK)
}
