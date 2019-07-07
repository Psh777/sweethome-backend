package sensor

import (
	"../../db/postgres"
	"../../types"
	"../../webserver/handlers"
	valid "github.com/asaskevich/govalidator"
	"net/http"
)

func GetDataSensor(w http.ResponseWriter, r *http.Request, sensorId string, sensorType string) {

	sensorTypeInt, err := valid.ToInt(sensorType)
	if err != nil {
		handlers.HandlerError(w, "sensor type invalid")
		return
	}

	if !valid.IsUUID(sensorId) {
		handlers.HandlerError(w, "sensor id invalid")
		return
	}

	data, err := postgres.GetData(sensorId, int(sensorTypeInt))
	if err != nil {
		handlers.HandlerError(w, err.Error())
		return
	}

	handlers.HandlerInterface(w, DataOut{Data: data})
}

type DataOut struct {
	Data []types.SensorData `json:"data"`
}
