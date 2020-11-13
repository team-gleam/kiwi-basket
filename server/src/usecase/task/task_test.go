package task

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/team-gleam/kiwi-basket/server/src/domain/model/task"
	"github.com/team-gleam/kiwi-basket/server/src/domain/model/user/credential"
	"github.com/team-gleam/kiwi-basket/server/src/domain/model/user/token"
	"github.com/team-gleam/kiwi-basket/server/src/domain/model/user/username"
	"github.com/team-gleam/kiwi-basket/server/src/domain/repository/mocks"
	credentialUsecase "github.com/team-gleam/kiwi-basket/server/src/usecase/user/credential"
)

func TestAdd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	credentialRepository := mocks.NewMockICredentialRepository(ctrl)
	loginRepository := mocks.NewMockILoginRepository(ctrl)
	taskRepository := mocks.NewMockITaskRepository(ctrl)

	usecase := NewTaskUsecase(
		credentialRepository,
		loginRepository,
		taskRepository,
	)

	t.Run("success", func(t *testing.T) {
		credentialRepository.EXPECT().Exists(gomock.Any()).Return(true, nil)

		username, _ := username.NewUsername("user")
		userToken := token.NewToken("123")
		auth := credential.NewAuth(username, userToken)
		credentialRepository.EXPECT().GetByToken(gomock.Any()).Return(auth, nil)

		taskRepository.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)

		task, _ := task.NewTask(0, "2020-10-10", "")
		err := usecase.Add(token.NewToken(""), task)
		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}
	})

	t.Run("has no credential", func(t *testing.T) {
		credentialRepository.EXPECT().Exists(gomock.Any()).Return(false, nil)

		task, _ := task.NewTask(0, "2020-10-10", "")
		err := usecase.Add(token.NewToken(""), task)
		if expected := credentialUsecase.InvalidToken; err.Error() != expected {
			t.Fatalf("expected: %v; got: %v\n", expected, err)
		}
	})
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	credentialRepository := mocks.NewMockICredentialRepository(ctrl)
	loginRepository := mocks.NewMockILoginRepository(ctrl)
	taskRepository := mocks.NewMockITaskRepository(ctrl)

	usecase := NewTaskUsecase(
		credentialRepository,
		loginRepository,
		taskRepository,
	)

	task1, _ := task.NewTask(1, "2020-01-01", "1")
	task2, _ := task.NewTask(2, "2020-01-01", "2")
	task3, _ := task.NewTask(3, "2020-01-01", "3")

	t.Run("success", func(t *testing.T) {
		credentialRepository.EXPECT().Exists(gomock.Any()).Return(true, nil)

		username, _ := username.NewUsername("user")
		userToken := token.NewToken("123")
		auth := credential.NewAuth(username, userToken)
		credentialRepository.EXPECT().GetByToken(gomock.Any()).Return(auth, nil)

		taskRepository.EXPECT().GetAll(gomock.Any()).Return(
			[]task.Task{task1, task2, task3},
			nil,
		)
		taskRepository.EXPECT().Remove(gomock.Any(), gomock.Any()).Return(nil)

		err := usecase.Delete(token.NewToken(""), 1)
		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}
	})

	t.Run("has no credential", func(t *testing.T) {
		credentialRepository.EXPECT().Exists(gomock.Any()).Return(false, nil)

		err := usecase.Delete(token.NewToken(""), 1)
		if expected := credentialUsecase.InvalidToken; err.Error() != expected {
			t.Fatalf("expected: %v; got: %v\n", expected, err)
		}
	})

	t.Run("with id 0", func(t *testing.T) {
		credentialRepository.EXPECT().Exists(gomock.Any()).Return(true, nil)

		err := usecase.Delete(token.NewToken(""), 0)
		if expected := IDIsNotZero; err.Error() != expected {
			t.Fatalf("expected: %v; got: %v\n", expected, err)
		}
	})

	t.Run("with negative id", func(t *testing.T) {
		credentialRepository.EXPECT().Exists(gomock.Any()).Return(true, nil)

		err := usecase.Delete(token.NewToken(""), -1)
		if expected := InvalidID; err.Error() != expected {
			t.Fatalf("expected: %v; got: %v\n", expected, err)
		}
	})

	t.Run("given invalid id", func(t *testing.T) {
		credentialRepository.EXPECT().Exists(gomock.Any()).Return(true, nil)

		username, _ := username.NewUsername("user")
		userToken := token.NewToken("123")
		auth := credential.NewAuth(username, userToken)
		credentialRepository.EXPECT().GetByToken(gomock.Any()).Return(auth, nil)

		taskRepository.EXPECT().GetAll(gomock.Any()).Return(
			[]task.Task{task1, task2, task3},
			nil,
		)

		err := usecase.Delete(token.NewToken(""), 4)
		if expected := InvalidID; err.Error() != expected {
			t.Fatalf("expected: %v; got: %v\n", expected, err)
		}
	})
}

func TestDeleteAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	credentialRepository := mocks.NewMockICredentialRepository(ctrl)
	loginRepository := mocks.NewMockILoginRepository(ctrl)
	taskRepository := mocks.NewMockITaskRepository(ctrl)

	usecase := NewTaskUsecase(
		credentialRepository,
		loginRepository,
		taskRepository,
	)

	t.Run("success", func(t *testing.T) {
		credentialRepository.EXPECT().Exists(gomock.Any()).Return(true, nil)

		username, _ := username.NewUsername("user")
		userToken := token.NewToken("123")
		auth := credential.NewAuth(username, userToken)
		credentialRepository.EXPECT().GetByToken(gomock.Any()).Return(auth, nil)

		taskRepository.EXPECT().RemoveAll(gomock.Any()).Return(nil)

		err := usecase.DeleteAll(token.NewToken(""))
		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}
	})

	t.Run("has no credential", func(t *testing.T) {
		credentialRepository.EXPECT().Exists(gomock.Any()).Return(false, nil)

		err := usecase.DeleteAll(token.NewToken(""))
		if expected := credentialUsecase.InvalidToken; err.Error() != expected {
			t.Fatalf("expected: %v; got: %v\n", expected, err)
		}
	})
}

func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	credentialRepository := mocks.NewMockICredentialRepository(ctrl)
	loginRepository := mocks.NewMockILoginRepository(ctrl)
	taskRepository := mocks.NewMockITaskRepository(ctrl)

	usecase := NewTaskUsecase(
		credentialRepository,
		loginRepository,
		taskRepository,
	)

	task1, _ := task.NewTask(1, "2020-01-01", "1")
	task2, _ := task.NewTask(2, "2020-01-01", "2")
	task3, _ := task.NewTask(3, "2020-01-01", "3")

	t.Run("success", func(t *testing.T) {
		credentialRepository.EXPECT().Exists(gomock.Any()).Return(true, nil)

		username, _ := username.NewUsername("user")
		userToken := token.NewToken("123")
		auth := credential.NewAuth(username, userToken)
		credentialRepository.EXPECT().GetByToken(gomock.Any()).Return(auth, nil)

		taskRepository.EXPECT().GetAll(gomock.Any()).Return(
			[]task.Task{task1, task2, task3},
			nil,
		)

		_, err := usecase.GetAll(token.NewToken(""))
		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}
	})

	t.Run("has no credential", func(t *testing.T) {
		credentialRepository.EXPECT().Exists(gomock.Any()).Return(false, nil)

		_, err := usecase.GetAll(token.NewToken(""))
		if expected := credentialUsecase.InvalidToken; err.Error() != expected {
			t.Fatalf("expected: %v; got: %v\n", expected, err)
		}
	})
}
