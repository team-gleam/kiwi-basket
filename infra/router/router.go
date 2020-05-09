package router

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/team-gleam/kiwi-basket/infra/db/handler"
	taskRepository "github.com/team-gleam/kiwi-basket/infra/db/task"
	timetablesRepository "github.com/team-gleam/kiwi-basket/infra/db/timetables"
	credentialRepository "github.com/team-gleam/kiwi-basket/infra/db/user/credential"
	loginRepository "github.com/team-gleam/kiwi-basket/infra/db/user/login"
	taskController "github.com/team-gleam/kiwi-basket/interfaces/controllers/task"
	timetablesController "github.com/team-gleam/kiwi-basket/interfaces/controllers/timetables"
	credentialController "github.com/team-gleam/kiwi-basket/interfaces/controllers/user/credential"
	loginController "github.com/team-gleam/kiwi-basket/interfaces/controllers/user/login"
)

func Run(c handler.Config) {
	e := echo.New()

	h, err := handler.NewDbHandler(c)
	if err != nil {
		log.Fatal(err)
	}

	taskRepo := taskRepository.NewTaskRepository(h)
	timetablesRepo := timetablesRepository.NewTimetablesRepository(h)
	credentialRepo := credentialRepository.NewCredentialRepository(h)
	loginRepo := loginRepository.NewLoginRepository(h)

	task := taskController.NewTaskController(
		credentialRepo,
		loginRepo,
		taskRepo,
	)

	timetables := timetablesController.NewTimetablesController(
		credentialRepo,
		loginRepo,
		timetablesRepo,
	)

	login := loginController.NewLoginController(
		loginRepo,
	)

	credential := credentialController.NewCredentialController(
		credentialRepo,
		loginRepo,
	)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/users", login.SignUp)
	e.DELETE("/users", login.DeleteAccound)

	e.POST("/tokens", credential.SignIn)

	e.POST("/timetables", timetables.Register)
	e.GET("/timetables", timetables.Get)

	e.POST("/tasks", task.Add)
	e.GET("/tasks", task.GetAll)
	e.DELETE("/tasks", task.Delete)

	e.Logger.Fatal(e.Start(":23450"))
}
