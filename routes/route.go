package routes

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitRouter(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/", index)

	authRouter(e)
	attendanceRouter(e)
	activityRouter(e)
}

func index(c echo.Context) error {
	return c.File("index.html")
}
