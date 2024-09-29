package utils

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func RequiresAuthToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("authorization")

		if auth != fmt.Sprintf("Bearer %s", AccessToken()) {
			return c.NoContent(401)
		}

		if err := next(c); err != nil {
			c.Error(err)
		}

		return nil
	}
}
