package postgres

import (
	"../../types"
	"fmt"
)

func NewData(sensor types.Sensor) error {

	row1 := DBX.QueryRow("UPDATE sensors SET update_timestamp = now(), alive = $1, request_id = $2 WHERE id = $3 RETURNING id;", sensor.Alive, sensor.RequestID, sensor.SensorID)
	var id string
	err := row1.Scan(&id)
	if err != nil {
		fmt.Println("NewData", err)
		return err
	}

	for i := 0; i < len(sensor.Data); i++ {
		_, err := DBX.Exec("INSERT INTO sensors_data (sensor_id, sensor_type, sensor_value, request_id) VALUES ($1, $2, $3, $4);", sensor.SensorID, sensor.Data[i].Type, sensor.Data[i].Value, sensor.RequestID)
		if err != nil {
			fmt.Println("NewData", err)
			return err
		}
	}
	return nil

}

func GetDataByType(sensorId string, sensorType int) ([]types.SensorData, error) {
	data := make([]types.SensorData, 0)
	err := DBX.Select(&data, "SELECT *, ROUND(extract(epoch from timestamp::timestamp with time zone) * 1000) as timestamp_int FROM sensors_data WHERE sensor_id = $1 AND sensor_type = $2 ORDER BY id DESC LIMIT 50;", sensorId, sensorType)
	if err != nil {
		fmt.Println("get Data: ", err)
		return nil, err
	}
	return data, nil
}

func GetDataByTypeActually(room string, sensorType int) (float64, error) {
	var sensorId string
	err := DBX.Get(&sensorId, "SELECT id FROM sensors WHERE room = $1 LIMIT 1;", room)
	if err != nil {
		fmt.Println("get SensorID: ", err)
		return 0, err
	}
	var data float64
	err = DBX.Get(&data, "SELECT sensor_value FROM sensors_data WHERE sensor_id = $1 AND sensor_type = $2 ORDER BY id DESC LIMIT 1;", sensorId, sensorType)
	if err != nil {
		fmt.Println("get Data: ", err)
		return 0, err
	}
	return data, nil
}

func GetDataByRequestID(requestID string) ([]types.SensorData, error) {
	data := make([]types.SensorData, 0)
	err := DBX.Select(&data, "SELECT *, ROUND(extract(epoch from timestamp::timestamp with time zone) * 1000) as timestamp_int FROM sensors_data WHERE request_id = $1 ORDER BY sensor_type DESC LIMIT 50;", requestID)
	if err != nil {
		fmt.Println("get Data: ", err)
		return nil, err
	}
	return data, nil
}


func GetSensors() ([]types.Sensor, error) {
	data := make([]types.Sensor, 0)
	err := DBX.Select(&data, "SELECT *, now() as time_now FROM sensors;")
	if err != nil {
		fmt.Println("get Sensors: ", err)
		return nil, err
	}
	return data, nil
}

func GetLastData(sensorId string) ([]types.SensorData, error) {
	sens := types.Sensor{}
	err := DBX.Get(&sens, "SELECT * FROM sensors WHERE id = $1;", sensorId)
	if err != nil {
		fmt.Println("get last data: ", err)
		return nil, err
	}
	data, err := GetDataByRequestID(sens.RequestID)
	if err != nil {
		fmt.Println("get last data: ", err)
		return nil, err
	}
	return data, nil
}

