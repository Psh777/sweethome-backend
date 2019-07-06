package sensor


import (
	"../../db/postgres"
	"../../webserver/handlers"
	"net/http"
)

func GetDataSensor(w http.ResponseWriter, r *http.Request, sensorId string) {
	data, err := postgres.GetData(sensorId)
	if err != nil {
		handlers.HandlerError(w, err.Error())
		return
	}
	handlers.HandlerInterface(w, data)
}