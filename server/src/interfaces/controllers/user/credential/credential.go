package credential

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/team-gleam/kiwi-basket/server/src/domain/model/user/token"
	credentialRepository "github.com/team-gleam/kiwi-basket/server/src/domain/repository/user/credential"
	loginRepository "github.com/team-gleam/kiwi-basket/server/src/domain/repository/user/login"
	errorResponse "github.com/team-gleam/kiwi-basket/server/src/interfaces/controllers/error"
	loginController "github.com/team-gleam/kiwi-basket/server/src/interfaces/controllers/user/login"
	credentialUsecase "github.com/team-gleam/kiwi-basket/server/src/usecase/user/credential"
)

type CredentialController struct {
	credentialUsecase credentialUsecase.CredentialUsecase
}

func NewCredentialController(
	c credentialRepository.ICredentialRepository,
	l loginRepository.ILoginRepository,
) *CredentialController {
	return &CredentialController{
		credentialUsecase.NewCredentialUsecase(c, l),
	}
}

type TokenResponse struct {
	Token string `json:"token"`
}

func toTokenResponse(t token.Token) TokenResponse {
	return TokenResponse{t.Token()}
}

func (c CredentialController) SignIn(ctx echo.Context) error {
	login := new(loginController.LoginResponse)
	err := ctx.Bind(login)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			errorResponse.NewError(fmt.Errorf(loginController.InvalidJSONFormat)),
		)
	}

	if !login.Validates() {
		return ctx.JSON(
			http.StatusUnauthorized,
			errorResponse.NewError(fmt.Errorf(loginController.InvalidUsernameOrPassword)),
		)
	}

	l, err := login.ToLogin()
	if err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			errorResponse.NewError(fmt.Errorf(errorResponse.InternalServerError)),
		)
	}

	token, err := c.credentialUsecase.Generate(l)
	if err != nil && err.Error() == credentialUsecase.UserNotFound {
		return ctx.JSON(
			http.StatusNotFound,
			errorResponse.NewError(err),
		)
	}
	if err != nil && err.Error() == credentialUsecase.InvalidUsernameOrPassword {
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

	return ctx.JSON(http.StatusOK, toTokenResponse(token))
}
