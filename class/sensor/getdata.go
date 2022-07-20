package sensor

import (
	"github.com/Psh777/sweethome-backend/db/postgres"
	"github.com/Psh777/sweethome-backend/types"
	"github.com/Psh777/sweethome-backend/webserver/handlers"
	valid "github.com/asaskevich/govalidator"
	"net/http"
	"sort"
)

func GetDataSensor(w http.ResponseWriter, _ *http.Request, sensorId string, sensorType string) {

	sensorTypeInt, err := valid.ToInt(sensorType)
	if err != nil {
		handlers.HandlerError(w, "sensor type invalid")
		return
	}

	if !valid.IsUUID(sensorId) {
		handlers.HandlerError(w, "sensor id invalid")
		return
	}

	data, err := postgres.GetDataByType(sensorId, int(sensorTypeInt))
	if err != nil {
		handlers.HandlerError(w, err.Error())
		return
	}

	sort.Sort(types.SensorDataByTime(data))

	lastData, err := postgres.GetLastData(sensorId)
	if err != nil {
		handlers.HandlerError(w, err.Error())
		return
	}

	sort.Sort(types.SensorDataByType(lastData))

	handlers.HandlerInterface(w, DataOut{Data: data, LastData: lastData, Actually: data[len(data)-1].Value})
}

type DataOut struct {
	Data     []types.SensorData `json:"data"`
	LastData []types.SensorData `json:"last_data"`
	Actually float64            `json:"actually"`
}
