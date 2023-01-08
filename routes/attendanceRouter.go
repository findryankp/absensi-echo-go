package routes

import (
	"absensi/controllers"
	"absensi/middlewares"

	"github.com/labstack/echo/v4"
)

func attendanceRouter(e *echo.Echo) {
	g := e.Group("/attendance")
	g.Use(middlewares.Authentication)
	g.POST("", controllers.CreateAttendance)
	g.GET("/riwayat", controllers.AttendaceRiwayat)
}
