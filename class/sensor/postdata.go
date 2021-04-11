package sensor

import (
	"../../db/postgres"
	"../../webserver/handlers"
	valid "github.com/asaskevich/govalidator"
	"net/http"
)

func PostDataSensor(w http.ResponseWriter, _ *http.Request, sensorId string, sensorType string) {

	sensorTypeInt, err := valid.ToInt(sensorType)
	if err != nil {
		handlers.HandlerError(w, "sensor type invalid")
		return
	}

	if !valid.IsUUID(sensorId) {
		handlers.HandlerError(w, "sensor id invalid")
		return
	}

	data, err := postgres.GetDataByTypeActually(sensorId, int(sensorTypeInt))
	if err != nil {
		handlers.HandlerError(w, err.Error())
		return
	}

	handlers.HandlerInterface(w, PostDataOut{Data: data})
}

type PostDataOut struct {
	Data float64 `json:"actually"`
}
