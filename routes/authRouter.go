package routes

import (
	"absensi/controllers"

	"github.com/labstack/echo/v4"
)

func authRouter(e *echo.Echo) {
	g := e.Group("/auth")
	g.POST("/register", controllers.Register)
	g.POST("/login", controllers.Login)
}
