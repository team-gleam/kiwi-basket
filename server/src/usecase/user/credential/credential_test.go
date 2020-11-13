package credential

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/team-gleam/kiwi-basket/server/src/domain/model/user/login"
	"github.com/team-gleam/kiwi-basket/server/src/domain/model/user/username"
	"github.com/team-gleam/kiwi-basket/server/src/domain/repository/mocks"
)

func TestGenerate(t *testing.T) {
	ctrl := gomock.NewController(t)
	credentialRepository := mocks.NewMockICredentialRepository(ctrl)
	loginRepository := mocks.NewMockILoginRepository(ctrl)
	usecase := NewCredentialUsecase(credentialRepository, loginRepository)

	t.Run("success", func(t *testing.T) {
		username, _ := username.NewUsername("user")
		password := "password"
		l := login.NewLogin(username, password)

		loginRepository.EXPECT().Exists(gomock.Any()).Return(true, nil)
		loginRepository.EXPECT().Get(gomock.Any()).Return(l, nil)

		credentialRepository.EXPECT().Remove(gomock.Any()).Return(nil)
		credentialRepository.EXPECT().Append(gomock.Any()).Return(nil)

		_, err := usecase.Generate(l)
		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}
	})

	t.Run("success", func(t *testing.T) {
		username, _ := username.NewUsername("user")
		password := "password"
		l := login.NewLogin(username, password)

		loginRepository.EXPECT().Exists(gomock.Any()).Return(true, nil)
		loginRepository.EXPECT().Get(gomock.Any()).Return(l, nil)

		credentialRepository.EXPECT().Remove(gomock.Any()).Return(nil)
		credentialRepository.EXPECT().Append(gomock.Any()).Return(nil)

		_, err := usecase.Generate(l)
		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}
	})

	t.Run("Verify return error", func(t *testing.T) {
		username, _ := username.NewUsername("user")
		password := "password"
		l := login.NewLogin(username, password)

		loginRepository.EXPECT().Exists(gomock.Any()).Return(true, fmt.Errorf("error occurred"))

		_, err := usecase.Generate(l)
		if err == nil {
			t.Fatalf("expected error but got nil")
		}
	})

	t.Run("not verified", func(t *testing.T) {
		username, _ := username.NewUsername("user")
		password := "password"
		l := login.NewLogin(username, password)

		loginRepository.EXPECT().Exists(gomock.Any()).Return(false, nil)

		_, err := usecase.Generate(l)
		if expected := InvalidUsernameOrPassword; err.Error() != expected {
			t.Fatalf("expected: %v; got: %v\n", expected, err)
		}
	})

	t.Run("Remove return error", func(t *testing.T) {
		username, _ := username.NewUsername("user")
		password := "password"
		l := login.NewLogin(username, password)

		loginRepository.EXPECT().Exists(gomock.Any()).Return(true, nil)
		loginRepository.EXPECT().Get(gomock.Any()).Return(l, nil)

		credentialRepository.EXPECT().Remove(gomock.Any()).Return(fmt.Errorf("error occurred"))

		_, err := usecase.Generate(l)
		if err == nil {
			t.Fatalf("expected error but got nil")
		}
	})

	t.Run("Append return error", func(t *testing.T) {
		username, _ := username.NewUsername("user")
		password := "password"
		l := login.NewLogin(username, password)

		loginRepository.EXPECT().Exists(gomock.Any()).Return(true, nil)
		loginRepository.EXPECT().Get(gomock.Any()).Return(l, nil)

		credentialRepository.EXPECT().Remove(gomock.Any()).Return(nil)
		credentialRepository.EXPECT().Append(gomock.Any()).Return(fmt.Errorf("error occurred"))

		_, err := usecase.Generate(l)
		if err == nil {
			t.Fatalf("expected error but got nil")
		}
	})
}
