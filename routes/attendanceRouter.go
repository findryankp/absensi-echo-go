package routes

import (
	"absensi/controllers"

	"github.com/labstack/echo/v4"
)

func attendanceRouter(e *echo.Echo) {
	g := e.Group("/attendance")
	g.POST("", controllers.CreateAttendance)
	g.GET("/riwayat", controllers.AttendaceRiwayat)
}
