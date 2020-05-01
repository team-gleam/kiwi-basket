package task

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	TaskModel "github.com/team-gleam/kiwi-basket/domain/model/task"
	"github.com/team-gleam/kiwi-basket/domain/model/user/token"
	taskRepository "github.com/team-gleam/kiwi-basket/domain/repository/task"
	credentialRepository "github.com/team-gleam/kiwi-basket/domain/repository/user/credential"
	loginRepository "github.com/team-gleam/kiwi-basket/domain/repository/user/login"
	errorResponse "github.com/team-gleam/kiwi-basket/interfaces/controllers/error"
	taskUsecase "github.com/team-gleam/kiwi-basket/usecase/task"
	credentialUsecase "github.com/team-gleam/kiwi-basket/usecase/user/credential"
)

type TaskController struct {
	taskUsecase taskUsecase.TaskUsecase
}

func NewTaskController(
	c credentialRepository.ICredentialRepository,
	l loginRepository.ILoginRepository,
	t taskRepository.ITaskRepository,
) *TaskController {
	return &TaskController{
		taskUsecase.NewTaskUsecase(
			credentialUsecase.NewCredentialUsecase(c, l),
			t,
		),
	}
}

const (
	InvalidJSONFormat = "invalid JSON format"
	InvalidID         = "invalid ID"
)

type TaskResponse struct {
	Task TaskJSON `json:"task"`
}

type TaskJSON struct {
	ID    string `json:"id"`
	Date  string `json:"date"`
	Title string `json:"title"`
}

func (t TaskResponse) toTask() (TaskModel.Task, error) {
	id, err := strconv.Atoi(t.Task.ID)
	if err != nil {
		return TaskModel.Task{}, err
	}
	if id == 0 {
		return TaskModel.Task{}, fmt.Errorf(InvalidID)
	}

	return TaskModel.NewTask(id, t.Task.Date, t.Task.Title)
}

func (c TaskController) Add(ctx echo.Context) error {
	t := ctx.Request().Header.Get("Token")
	if t == "" {
		return ctx.JSON(
			http.StatusUnauthorized,
			errorResponse.NewError(fmt.Errorf(credentialUsecase.InvalidToken)),
		)
	}

	res := new(TaskResponse)
	err := ctx.Bind(res)
	if err != nil || res.Task.ID == "" {
		return ctx.JSON(
			http.StatusBadRequest,
			errorResponse.NewError(fmt.Errorf(InvalidJSONFormat)),
		)
	}

	task, err := res.toTask()
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			errorResponse.NewError(fmt.Errorf(InvalidJSONFormat)),
		)
	}

	err = c.taskUsecase.Add(token.NewToken(t), task)
	if err.Error() == credentialUsecase.InvalidToken {
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

	return ctx.NoContent(http.StatusOK)
}

type ID struct {
	ID string `json:"id"`
}

func (c TaskController) Delete(ctx echo.Context) error {
	t := ctx.Request().Header.Get("Token")
	if t == "" {
		return ctx.JSON(
			http.StatusUnauthorized,
			errorResponse.NewError(fmt.Errorf(credentialUsecase.InvalidToken)),
		)
	}

	res := new(ID)
	err := ctx.Bind(res)
	if err != nil || res.ID == "" {
		return ctx.JSON(
			http.StatusBadRequest,
			errorResponse.NewError(fmt.Errorf(InvalidJSONFormat)),
		)
	}

	id, err := strconv.Atoi(res.ID)
	if err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			errorResponse.NewError(fmt.Errorf(errorResponse.InternalServerError)),
		)
	}

	err = c.taskUsecase.Delete(token.NewToken(t), id)
	if err.Error() == credentialUsecase.InvalidToken {
		return ctx.JSON(
			http.StatusUnauthorized,
			errorResponse.NewError(fmt.Errorf(credentialUsecase.InvalidToken)),
		)
	}
	if err != nil {
		return ctx.JSON(
			http.StatusInternalServerError,
			errorResponse.NewError(fmt.Errorf(errorResponse.InternalServerError)),
		)
	}

	return ctx.NoContent(http.StatusOK)
}

type TasksResponse struct {
	Tasks []TaskJSON `json:"tasks"`
}

func toTasksResponse(ts []TaskModel.Task) TasksResponse {
	res := []TaskJSON{}
	for _, t := range ts {
		res = append(res, TaskJSON{
			ID:    strconv.Itoa(t.ID()),
			Date:  t.TextDate(),
			Title: t.Title(),
		})
	}

	return TasksResponse{res}
}

func (c TaskController) GetAll(ctx echo.Context) error {
	t := ctx.Request().Header.Get("Token")
	if t == "" {
		return ctx.JSON(
			http.StatusUnauthorized,
			errorResponse.NewError(fmt.Errorf(credentialUsecase.InvalidToken)),
		)
	}

	tasks, err := c.taskUsecase.GetAll(token.NewToken(t))
	if err.Error() == credentialUsecase.InvalidToken {
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

	return ctx.JSON(http.StatusOK, toTasksResponse(tasks))
}
