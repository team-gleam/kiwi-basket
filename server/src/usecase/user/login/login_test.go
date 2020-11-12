package login

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/team-gleam/kiwi-basket/server/src/domain/model/user/login"
	"github.com/team-gleam/kiwi-basket/server/src/domain/model/user/username"
	"github.com/team-gleam/kiwi-basket/server/src/domain/repository/mocks"
)

func TestAdd(t *testing.T) {
	ctrl := gomock.NewController(t)
	loginRepository := mocks.NewMockILoginRepository(ctrl)
	usecase := NewLoginUsecase(loginRepository)

	t.Run("success", func(t *testing.T) {
		loginRepository.EXPECT().Exists(gomock.Any()).Return(false, nil)
		loginRepository.EXPECT().Create(gomock.Any()).Return(nil)

		username, _ := username.NewUsername("user")
		err := usecase.Add(login.NewLogin(username, ""))
		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}
	})

	t.Run("user already exists", func(t *testing.T) {
		loginRepository.EXPECT().Exists(gomock.Any()).Return(true, nil)

		username, _ := username.NewUsername("user")
		err := usecase.Add(login.NewLogin(username, ""))
		if err.Error() != UsernameAlreadyExists {
			t.Fatalf("expected: %v; got: %v\n", UsernameAlreadyExists, err)
		}
	})

	t.Run("Exists return error", func(t *testing.T) {
		loginRepository.EXPECT().Exists(gomock.Any()).Return(false, fmt.Errorf("error occurred"))

		username, _ := username.NewUsername("user")
		err := usecase.Add(login.NewLogin(username, ""))
		if err == nil {
			t.Fatalf("unexpected error but got nil")
		}
	})

	t.Run("Create return error", func(t *testing.T) {
		loginRepository.EXPECT().Exists(gomock.Any()).Return(false, nil)
		loginRepository.EXPECT().Create(gomock.Any()).Return(fmt.Errorf("error occurred"))

		username, _ := username.NewUsername("user")
		err := usecase.Add(login.NewLogin(username, ""))
		if err == nil {
			t.Fatalf("unexpected error but got nil")
		}
	})
}
