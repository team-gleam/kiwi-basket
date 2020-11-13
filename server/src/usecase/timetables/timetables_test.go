package timetables

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/team-gleam/kiwi-basket/server/src/domain/model/timetables"
	"github.com/team-gleam/kiwi-basket/server/src/domain/model/user/credential"
	"github.com/team-gleam/kiwi-basket/server/src/domain/model/user/token"
	"github.com/team-gleam/kiwi-basket/server/src/domain/model/user/username"
	"github.com/team-gleam/kiwi-basket/server/src/domain/repository/mocks"
)

var (
	ts = timetables.NewTimetables(
		timetables.NewTimetable(
			timetables.NoRoom("11", ""),
			timetables.NoRoom("12", ""),
			timetables.NoRoom("13", ""),
			timetables.NoRoom("14", ""),
			timetables.NoRoom("15", ""),
		),
		timetables.NewTimetable(
			timetables.NoRoom("21", ""),
			timetables.NoRoom("22", ""),
			timetables.NoRoom("23", ""),
			timetables.NoRoom("24", ""),
			timetables.NoRoom("25", ""),
		),
		timetables.NewTimetable(
			timetables.NoRoom("31", ""),
			timetables.NoRoom("32", ""),
			timetables.NoRoom("33", ""),
			timetables.NoRoom("34", ""),
			timetables.NoRoom("35", ""),
		),
		timetables.NewTimetable(
			timetables.NoRoom("41", ""),
			timetables.NoRoom("42", ""),
			timetables.NoRoom("43", ""),
			timetables.NoRoom("44", ""),
			timetables.NoRoom("45", ""),
		),
		timetables.NewTimetable(
			timetables.NoRoom("51", ""),
			timetables.NoRoom("52", ""),
			timetables.NoRoom("53", ""),
			timetables.NoRoom("54", ""),
			timetables.NoRoom("55", ""),
		),
	)
)

func TestAdd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	credentialRepository := mocks.NewMockICredentialRepository(ctrl)
	loginRepository := mocks.NewMockILoginRepository(ctrl)
	timetablesRepository := mocks.NewMockITimetablesRepository(ctrl)

	usecase := NewTimetablesUsecase(
		credentialRepository,
		loginRepository,
		timetablesRepository,
	)

	t.Run("success", func(t *testing.T) {
		credentialRepository.EXPECT().Exists(gomock.Any()).Return(true, nil)

		username, _ := username.NewUsername("user")
		userToken := token.NewToken("123")
		auth := credential.NewAuth(username, userToken)
		credentialRepository.EXPECT().GetByToken(gomock.Any()).Return(auth, nil)

		timetablesRepository.EXPECT().Exists(gomock.Any()).Return(false, nil)
		timetablesRepository.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)

		err := usecase.Add(token.NewToken(""), ts)
		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}
	})

	t.Run("has no credential", func(t *testing.T) {
		credentialRepository.EXPECT().Exists(gomock.Any()).Return(false, nil)

		err := usecase.Add(token.NewToken(""), ts)
		if err == nil {
			t.Fatalf("expected error but got nil")
		}
	})
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	credentialRepository := mocks.NewMockICredentialRepository(ctrl)
	loginRepository := mocks.NewMockILoginRepository(ctrl)
	timetablesRepository := mocks.NewMockITimetablesRepository(ctrl)

	usecase := NewTimetablesUsecase(
		credentialRepository,
		loginRepository,
		timetablesRepository,
	)

	t.Run("success", func(t *testing.T) {
		credentialRepository.EXPECT().Exists(gomock.Any()).Return(true, nil)

		username, _ := username.NewUsername("user")
		userToken := token.NewToken("123")
		auth := credential.NewAuth(username, userToken)
		credentialRepository.EXPECT().GetByToken(gomock.Any()).Return(auth, nil)

		timetablesRepository.EXPECT().Exists(gomock.Any()).Return(true, nil)
		timetablesRepository.EXPECT().Delete(gomock.Any()).Return(nil)

		err := usecase.Delete(userToken)
		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}
	})

	t.Run("has no credential", func(t *testing.T) {
		credentialRepository.EXPECT().Exists(gomock.Any()).Return(false, nil)

		err := usecase.Add(token.NewToken(""), ts)
		if err == nil {
			t.Fatalf("expected error but got nil")
		}
	})
}

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	credentialRepository := mocks.NewMockICredentialRepository(ctrl)
	loginRepository := mocks.NewMockILoginRepository(ctrl)
	timetablesRepository := mocks.NewMockITimetablesRepository(ctrl)

	usecase := NewTimetablesUsecase(
		credentialRepository,
		loginRepository,
		timetablesRepository,
	)

	t.Run("success", func(t *testing.T) {
		credentialRepository.EXPECT().Exists(gomock.Any()).Return(true, nil)

		username, _ := username.NewUsername("user")
		userToken := token.NewToken("123")
		auth := credential.NewAuth(username, userToken)
		credentialRepository.EXPECT().GetByToken(gomock.Any()).Return(auth, nil)

		timetablesRepository.EXPECT().Exists(gomock.Any()).Return(true, nil)
		timetablesRepository.EXPECT().Get(gomock.Any()).Return(ts, nil)

		_, err := usecase.Get(userToken)
		if err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}
	})

	t.Run("has no credential", func(t *testing.T) {
		credentialRepository.EXPECT().Exists(gomock.Any()).Return(false, nil)

		err := usecase.Add(token.NewToken(""), ts)
		if err == nil {
			t.Fatalf("expected error but got nil")
		}
	})

	t.Run("timetables not found", func(t *testing.T) {
		credentialRepository.EXPECT().Exists(gomock.Any()).Return(true, nil)

		username, _ := username.NewUsername("user")
		userToken := token.NewToken("123")
		auth := credential.NewAuth(username, userToken)
		credentialRepository.EXPECT().GetByToken(gomock.Any()).Return(auth, nil)

		timetablesRepository.EXPECT().Exists(gomock.Any()).Return(false, nil)

		_, err := usecase.Get(userToken)
		if expected := TimetablesNotFound; err.Error() != expected {
			t.Fatalf("expected: %v; got: %v\n", expected, err)
		}
	})
}
