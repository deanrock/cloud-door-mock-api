package routes

import (
	"github.com/deanrock/cloud-door-mock-api/utils"
	"github.com/labstack/echo/v4"
)

func InitTokenRoutes(e *echo.Echo) {
	e.POST("/token", func(c echo.Context) error {
		client_id := c.FormValue("client_id")
		grant_type := c.FormValue("password")
		username := c.FormValue("username")
		password := c.FormValue("password")

		if !utils.IsFormEncoded(c.Request()) {
			return c.NoContent(500)
		}

		if client_id == "DoorCloudWebApp" && grant_type == "password" && username == "user@example.com" && password == "password" {
			data := struct {
				Expires      string `json:".expires"`
				Issued       string `json:".issued"`
				AccessToken  string `json:"access_token"`
				AsClientId   string `json:"as:client_id"`
				ExpiresIn    int    `json:"expires_in"`
				RefreshToken string `json:"refresh_token"`
				TokenType    string `json:"token_type"`
				UserName     string `json:"userName"`
			}{
				Expires:      "Sun, 29 Sep 2024 13:38:55 GMT",
				Issued:       "Sun, 29 Sep 2024 12:38:55 GMT",
				AccessToken:  utils.AccessToken(),
				AsClientId:   "DoorCloudWebApp",
				ExpiresIn:    3599,
				RefreshToken: "gwx4g2bi2ydu3wg8eg6p5dyedtfk53ag",
				TokenType:    "bearer",
				UserName:     "user@example.com",
			}

			return c.JSON(200, data)
		}

		return c.NoContent(500)
	})
}
