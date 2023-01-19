package app

import (
	"github.com/Levap123/test-task-days/internal/service"
	"github.com/Levap123/test-task-days/internal/transport/rest"
	"github.com/labstack/echo/v4"
)

type App struct {
	handler *rest.Handler
	service *service.Service
	echo    *echo.Echo
}

func NewApp() *App {
	app := new(App)
	app.service = service.NewService()
	app.handler = rest.NewHandler(app.service)
	app.echo = app.handler.InitRoutes()
	return app
}

func (app *App) Run() error {
	return app.echo.Start(":8080")
}
