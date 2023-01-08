package middlewares

import (
	"absensi/controllers"
	"absensi/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if isExistAttendance := controllers.IsExistAttendance(c, 1); isExistAttendance {
			return helpers.ResponseJson(c, http.StatusUnauthorized, false, "You haven't been checkin today", nil)
		}
		return next(c)
	}
}
