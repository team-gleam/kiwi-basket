package timetables

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/team-gleam/kiwi-basket/server/src/domain/model/timetables"
	"github.com/team-gleam/kiwi-basket/server/src/domain/model/user/credential"
	"github.com/team-gleam/kiwi-basket/server/src/domain/model/user/token"
	"github.com/team-gleam/kiwi-basket/server/src/domain/model/user/username"
	"github.com/team-gleam/kiwi-basket/server/src/domain/repository/mocks"
	credentialUsecase "github.com/team-gleam/kiwi-basket/server/src/usecase/user/credential"
)

var (
	ts = timetables.NewTimetables(
		timetables.NewTimetable(
			timetables.NoRoom("11", "a"),
			timetables.NoRoom("12", "a"),
			timetables.NoRoom("13", "a"),
			timetables.NoRoom("14", "a"),
			timetables.NoRoom("15", "a"),
		),
		timetables.NewTimetable(
			timetables.NoRoom("21", "b"),
			timetables.NoRoom("22", "b"),
			timetables.NoRoom("23", "b"),
			timetables.NoRoom("24", "b"),
			timetables.NoRoom("25", "b"),
		),
		timetables.NewTimetable(
			timetables.NoRoom("31", "c"),
			timetables.NoRoom("32", "c"),
			timetables.NoRoom("33", "c"),
			timetables.NoRoom("34", "c"),
			timetables.NoRoom("35", "c"),
		),
		timetables.NewTimetable(
			timetables.NoRoom("41", "d"),
			timetables.NoRoom("42", "d"),
			timetables.NoRoom("43", "d"),
			timetables.NoRoom("44", "d"),
			timetables.NoRoom("45", "d"),
		),
		timetables.NewTimetable(
			timetables.NoRoom("51", "e"),
			timetables.NoRoom("52", "e"),
			timetables.NoRoom("53", "e"),
			timetables.NoRoom("54", "e"),
			timetables.NoRoom("55", "e"),
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

		err := usecase.Delete(token.NewToken(""))
		if expected := credentialUsecase.InvalidToken; err.Error() != expected {
			t.Fatalf("expected: %v; got: %v\n", expected, err)
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

		_, err := usecase.Get(token.NewToken("a"))
		if expected := credentialUsecase.InvalidToken; err.Error() != expected {
			t.Fatalf("expected: %v; got: %v\n", expected, err)
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
