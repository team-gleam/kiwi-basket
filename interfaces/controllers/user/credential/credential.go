package credential

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	credentialRepository "github.com/team-gleam/kiwi-basket/domain/repository/user/credential"
	loginRepository "github.com/team-gleam/kiwi-basket/domain/repository/user/login"
	errorResponse "github.com/team-gleam/kiwi-basket/interfaces/controllers/error"
	loginController "github.com/team-gleam/kiwi-basket/interfaces/controllers/user/login"
	credentialUsecase "github.com/team-gleam/kiwi-basket/usecase/user/credential"
)

type CredentialController struct {
	credentialUsecase credentialUsecase.CredentialUsecase
}

func NewCredentialController(
	c credentialRepository.ICredentialRepository,
	l loginRepository.ILoginRepository,
) *CredentialController {
	return &CredentialController{
		credentialUsecase.NewCredentialRepository(c, l),
	}
}

type Token struct {
	Token string `json:"token"`
}

func NewToken(t string) Token {
	return Token{t}
}

func (c CredentialController) SignIn(ctx echo.Context) error {
	login := new(loginController.LoginResponse)
	err := ctx.Bind(login)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			errorResponse.NewError(fmt.Errorf(loginController.InvalidJsonFormat)),
		)
	}
	if login.Username == "" || login.Password == "" {
		return ctx.JSON(
			http.StatusUnauthorized,
			errorResponse.NewError(fmt.Errorf(loginController.InvalidUsernameOrPassword)),
		)
	}

	l, err := login.ToLogin()
	if err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			errorResponse.NewError(fmt.Errorf(loginController.InternalServerError)),
		)
	}

	token, err := c.credentialUsecase.Generate(l)
	if err.Error() == credentialUsecase.UserNotFound {
		return ctx.JSON(
			http.StatusNotFound,
			errorResponse.NewError(fmt.Errorf(credentialUsecase.UserNotFound)),
		)
	}
	if err.Error() == credentialUsecase.InvalidUsernameOrPassword {
		return ctx.JSON(
			http.StatusUnauthorized,
			errorResponse.NewError(fmt.Errorf(credentialUsecase.InvalidUsernameOrPassword)),
		)
	}
	if err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			errorResponse.NewError(fmt.Errorf(loginController.InternalServerError)),
		)
	}

	return ctx.JSON(http.StatusOK, NewToken(token.Token()))
}
