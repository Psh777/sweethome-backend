package sensor

import (
	"../../db/postgres"
	"../../types"
	"../../webserver/handlers"
	"net/http"
)

func GetSensors(w http.ResponseWriter, _ *http.Request) {
	sensors, err := postgres.GetSensors()
	if err != nil {
		handlers.HandlerError(w, err.Error())
		return
	}
	handlers.HandlerInterface(w, SensorsOut{
		Sensors: sensors,
	})
}

type SensorsOut struct {
	Sensors []types.Sensor `json:"sensors"`
}
