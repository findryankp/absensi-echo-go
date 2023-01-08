package routes

import (
	"absensi/controllers"

	"github.com/labstack/echo/v4"
)

func activityRouter(e *echo.Echo) {
	g := e.Group("/activity")
	g.GET("", controllers.GetAllActivity)
	g.GET("/:id", controllers.GetActivity)
	g.POST("", controllers.CreateActivity)
	g.PUT("/:id", controllers.UpdateActivity)
	g.DELETE("/:id", controllers.DeleteActivity)
	g.GET("/riwayat", controllers.GetAllActivity)
}
