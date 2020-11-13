package credential

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/team-gleam/kiwi-basket/server/src/domain/model/user/credential"
	"github.com/team-gleam/kiwi-basket/server/src/domain/model/user/login"
	"github.com/team-gleam/kiwi-basket/server/src/domain/model/user/token"
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

func TestDelete(t *testing.T) {
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

		err := usecase.Delete(l)
		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}
	})

	t.Run("Verify return error", func(t *testing.T) {
		username, _ := username.NewUsername("user")
		password := "password"
		l := login.NewLogin(username, password)

		loginRepository.EXPECT().Exists(gomock.Any()).Return(true, fmt.Errorf("error occurred"))

		err := usecase.Delete(l)
		if err == nil {
			t.Fatalf("expected error but got nil")
		}
	})

	t.Run("user not found", func(t *testing.T) {
		username, _ := username.NewUsername("user")
		password := "password"
		l := login.NewLogin(username, password)

		loginRepository.EXPECT().Exists(gomock.Any()).Return(false, nil)

		err := usecase.Delete(l)
		if expected := "username not found"; err.Error() != expected {
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

		err := usecase.Delete(l)
		if err == nil {
			t.Fatalf("expected error but got nil")
		}
	})
}

func TestGet(t *testing.T) {
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

		token := token.NewToken("123")
		auth := credential.NewAuth(username, token)
		credentialRepository.EXPECT().GetByUsername(gomock.Any()).Return(auth, nil)

		v, err := usecase.Get(l)
		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}
		if v != auth {
			t.Fatalf("expected: %v; got: %v\n", auth, v)
		}
	})

	t.Run("Verify return error", func(t *testing.T) {
		username, _ := username.NewUsername("user")
		password := "password"
		l := login.NewLogin(username, password)

		loginRepository.EXPECT().Exists(gomock.Any()).Return(true, fmt.Errorf("error occurred"))

		_, err := usecase.Get(l)
		if err == nil {
			t.Fatalf("expected error but got nil")
		}
	})

	t.Run("username not found", func(t *testing.T) {
		username, _ := username.NewUsername("user")
		password := "password"
		l := login.NewLogin(username, password)

		loginRepository.EXPECT().Exists(gomock.Any()).Return(false, nil)

		_, err := usecase.Get(l)
		if expected := "username not found"; err.Error() != expected {
			t.Fatalf("expected: %v; got: %v\n", expected, err)
		}
	})

	t.Run("GetByUsername return error", func(t *testing.T) {
		username, _ := username.NewUsername("user")
		password := "password"
		l := login.NewLogin(username, password)

		loginRepository.EXPECT().Exists(gomock.Any()).Return(true, nil)
		loginRepository.EXPECT().Get(gomock.Any()).Return(l, nil)

		token := token.NewToken("123")
		auth := credential.NewAuth(username, token)
		credentialRepository.EXPECT().GetByUsername(gomock.Any()).Return(auth, fmt.Errorf("error occurred"))

		_, err := usecase.Get(l)
		if err == nil {
			t.Fatalf("expected error but got nil")
		}
	})
}

func TestHasCredential(t *testing.T) {
	ctrl := gomock.NewController(t)
	credentialRepository := mocks.NewMockICredentialRepository(ctrl)
	loginRepository := mocks.NewMockILoginRepository(ctrl)
	usecase := NewCredentialUsecase(credentialRepository, loginRepository)

	t.Run("success", func(t *testing.T) {
		credentialRepository.EXPECT().Exists(gomock.Any()).Return(true, nil)

		v, err := usecase.HasCredential(token.NewToken(""))
		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}
		if !v {
			t.Fatalf("expected: %v; got: %v\n", true, v)
		}
	})

	t.Run("has not credential", func(t *testing.T) {
		credentialRepository.EXPECT().Exists(gomock.Any()).Return(false, nil)

		v, err := usecase.HasCredential(token.NewToken(""))
		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}
		if v {
			t.Fatalf("expected: %v; got: %v\n", false, v)
		}
	})
}

func TestWhose(t *testing.T) {
	ctrl := gomock.NewController(t)
	credentialRepository := mocks.NewMockICredentialRepository(ctrl)
	loginRepository := mocks.NewMockILoginRepository(ctrl)
	usecase := NewCredentialUsecase(credentialRepository, loginRepository)

	t.Run("success", func(t *testing.T) {
		username, _ := username.NewUsername("user")
		token := token.NewToken("123")
		auth := credential.NewAuth(username, token)

		credentialRepository.EXPECT().GetByToken(gomock.Any()).Return(auth, nil)

		u, err := usecase.Whose(token)
		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}
		if u != username {
			t.Fatalf("expected: %v; got: %v\n", username, u)
		}
	})

	t.Run("GetByToken return error", func(t *testing.T) {
		username, _ := username.NewUsername("user")
		token := token.NewToken("123")
		auth := credential.NewAuth(username, token)

		credentialRepository.EXPECT().GetByToken(gomock.Any()).Return(auth, fmt.Errorf("error occurred"))

		_, err := usecase.Whose(token)
		if err == nil {
			t.Fatalf("expected error but got nil")
		}
	})
}
