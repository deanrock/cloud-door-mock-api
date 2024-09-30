package routes

import (
	"github.com/deanrock/cloud-door-mock-api/client"
	"github.com/deanrock/cloud-door-mock-api/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func InitLocationRoutes(e *echo.Echo) {
	doors := []struct {
		Name string
		Id   uuid.UUID
	}{{
		Name: "Floor #1",
		Id:   uuid.MustParse("de75458c-7bbb-4b33-b5d8-dc69b565357a"),
	}, {
		Name: "Floor #2",
		Id:   uuid.MustParse("e77afeb3-bbdd-44bb-8281-e9b483ea2664"),
	}, {
		Name: "Floor #3",
		Id:   uuid.MustParse("3ce2c622-2d6f-4755-8a06-9d91f27d56bc"),
	}}

	location := e.Group("/api/Location")
	location.Use(utils.RequiresAuthToken)

	location.GET("/GetUserLocations", func(c echo.Context) error {
		var data []client.DoorCloudLocationDtosUserLocationDto
		for _, door := range doors {
			data = append(data, client.DoorCloudLocationDtosUserLocationDto{
				Beacons:          utils.Pointer([]client.DoorCloudBeaconDtosBeaconDto{}),
				ControllerOnline: utils.Pointer(true),
				Geolocations: utils.Pointer([]client.DoorCloudGeolocationDtosGeolocationDto{{
					Id:        utils.Pointer(uuid.New()),
					Latitude:  utils.Pointer(45.64993864503844),
					Longitude: utils.Pointer(13.775274149794846),
					Name:      utils.Pointer("Geolocation #X"),
				}}),
				Id:         utils.Pointer(door.Id),
				IsFavorite: utils.Pointer(false),
				Name:       utils.Pointer(door.Name),
				OutputId:   utils.Pointer(uuid.New()),
				OutputName: utils.Pointer("Door #X Output"),
			})
		}
		return c.JSON(200, utils.ToAbpResponse(data))
	})

	location.POST("/OpenDoorOnLocation", func(c echo.Context) error {
		if !utils.IsFormEncoded(c.Request()) {
			return c.NoContent(500)
		}

		accessPointId := c.FormValue("accessPointId")

		for _, door := range doors {
			if door.Id.String() == accessPointId {
				data := client.DoorCloudCommandLogDtosCommandLogReference{
					Id: utils.Pointer(door.Id),
				}
				return c.JSON(200, utils.ToAbpResponse(data))
			}
		}

		return c.NoContent(400)
	})
}
