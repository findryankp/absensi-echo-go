package middlewares

import (
	"absensi/helpers"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, tokenString := helpers.HeaderToken(c)

		if token := c.Request().Header.Get(os.Getenv("HEADERAUTH")); tokenString == "" || tokenString == token {
			return helpers.ResponseJson(c, http.StatusUnauthorized, false, "request does not contain an access token", nil)
		}

		if err := helpers.ValidateToken(tokenString); err != nil {
			return helpers.ResponseJson(c, http.StatusUnauthorized, false, "request does not contain an valid token", nil)
		}
		return next(c)
	}
}
