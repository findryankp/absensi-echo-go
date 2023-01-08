package routes

import "github.com/labstack/echo/v4"

func InitRouter(e *echo.Echo) {
	authRouter(e)
	attendanceRouter(e)
	activityRouter(e)
}
