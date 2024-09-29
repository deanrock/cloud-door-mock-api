package main

import (
	"github.com/deanrock/cloud-door-mock-api/routes"
	"github.com/labstack/echo/v4"
)

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=resources/config.yaml resources/openapi.yaml

func main() {
	e := echo.New()
	routes.InitTokenRoutes(e)
	routes.InitAdminRoutes(e)
	routes.InitLocationRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
