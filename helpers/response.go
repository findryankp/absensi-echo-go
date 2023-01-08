package helpers

import "github.com/labstack/echo/v4"

func ResponseJson(c echo.Context, code int, status, msg, data interface{}) error {
	return c.JSON(code, map[string]interface{}{
		"status":  status,
		"message": msg,
		"data":    data,
	})
}
