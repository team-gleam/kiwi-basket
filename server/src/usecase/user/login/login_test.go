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
		if expected := UsernameAlreadyExists; err.Error() != expected {
			t.Fatalf("expected: %v; got: %v\n", expected, err)
		}
	})

	t.Run("Exists return error", func(t *testing.T) {
		loginRepository.EXPECT().Exists(gomock.Any()).Return(false, fmt.Errorf("error occurred"))

		username, _ := username.NewUsername("user")
		err := usecase.Add(login.NewLogin(username, ""))
		if err == nil {
			t.Fatalf("expected error but got nil")
		}
	})

	t.Run("Create return error", func(t *testing.T) {
		loginRepository.EXPECT().Exists(gomock.Any()).Return(false, nil)
		loginRepository.EXPECT().Create(gomock.Any()).Return(fmt.Errorf("error occurred"))

		username, _ := username.NewUsername("user")
		err := usecase.Add(login.NewLogin(username, ""))
		if err == nil {
			t.Fatalf("expected error but got nil")
		}
	})
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	loginRepository := mocks.NewMockILoginRepository(ctrl)
	usecase := NewLoginUsecase(loginRepository)

	t.Run("success", func(t *testing.T) {
		loginRepository.EXPECT().Exists(gomock.Any()).Return(true, nil)
		loginRepository.EXPECT().Delete(gomock.Any()).Return(nil)

		username, _ := username.NewUsername("user")
		err := usecase.Delete(login.NewLogin(username, ""))
		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}
	})

	t.Run("user not found", func(t *testing.T) {
		loginRepository.EXPECT().Exists(gomock.Any()).Return(false, nil)

		username, _ := username.NewUsername("user")
		err := usecase.Delete(login.NewLogin(username, ""))
		if expected := UsernameNotFound; err.Error() != expected {
			t.Fatalf("expected: %v; got: %v\n", expected, err)
		}
	})

	t.Run("Exists return error", func(t *testing.T) {
		loginRepository.EXPECT().Exists(gomock.Any()).Return(true, fmt.Errorf("error occurred"))

		username, _ := username.NewUsername("user")
		err := usecase.Delete(login.NewLogin(username, ""))
		if err == nil {
			t.Fatalf("expected error but got nil")
		}
	})

	t.Run("Delete return error", func(t *testing.T) {
		loginRepository.EXPECT().Exists(gomock.Any()).Return(true, nil)
		loginRepository.EXPECT().Delete(gomock.Any()).Return(fmt.Errorf("error occurred"))

		username, _ := username.NewUsername("user")
		err := usecase.Delete(login.NewLogin(username, ""))
		if err == nil {
			t.Fatalf("expected error but got nil")
		}
	})
}

func TestVerify(t *testing.T) {
	ctrl := gomock.NewController(t)
	loginRepository := mocks.NewMockILoginRepository(ctrl)
	usecase := NewLoginUsecase(loginRepository)

	t.Run("success", func(t *testing.T) {
		username, _ := username.NewUsername("user")
		password := "password"
		login := login.NewLogin(username, password)

		loginRepository.EXPECT().Exists(gomock.Any()).Return(true, nil)
		loginRepository.EXPECT().Get(gomock.Any()).Return(login, nil)

		v, err := usecase.Verify(login)
		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}
		if !v {
			t.Fatalf("expected: %v; got: %v\n", true, err)
		}
	})

	t.Run("user not found", func(t *testing.T) {
		username, _ := username.NewUsername("user")
		password := "password"
		login := login.NewLogin(username, password)

		loginRepository.EXPECT().Exists(gomock.Any()).Return(false, nil)

		_, err := usecase.Verify(login)
		if expected := UsernameNotFound; err.Error() != expected {
			t.Fatalf("expected: %v; got: %v\n", expected, err)
		}
	})

	t.Run("invalid password", func(t *testing.T) {
		username, _ := username.NewUsername("user")
		password := "password1"
		l := login.NewLogin(username, password)

		loginRepository.EXPECT().Exists(gomock.Any()).Return(true, nil)
		loginRepository.EXPECT().Get(gomock.Any()).Return(login.NewLogin(username, "password2"), nil)

		v, err := usecase.Verify(l)
		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}
		if v {
			t.Fatalf("expected: %v; got: %v\n", false, err)
		}
	})

	t.Run("Exists return error", func(t *testing.T) {
		username, _ := username.NewUsername("user")
		password := "password"
		l := login.NewLogin(username, password)

		loginRepository.EXPECT().Exists(gomock.Any()).Return(true, fmt.Errorf("error occurred"))

		_, err := usecase.Verify(l)
		if err == nil {
			t.Fatalf("expected error but got nil")
		}
	})

	t.Run("Get return error", func(t *testing.T) {
		username, _ := username.NewUsername("user")
		password := "password"
		l := login.NewLogin(username, password)

		loginRepository.EXPECT().Exists(gomock.Any()).Return(true, nil)
		loginRepository.EXPECT().Get(gomock.Any()).Return(l, fmt.Errorf("error occurred"))

		_, err := usecase.Verify(l)
		if err == nil {
			t.Fatalf("expected error but got nil")
		}
	})
}
