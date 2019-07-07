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

func GetData(sensorId string, sensorType int) ([]types.SensorData, error) {
	fmt.Println("get data: ", sensorId)
	data := make([]types.SensorData, 0)
	err := DBX.Select(&data, "SELECT *, ROUND(extract(epoch from timestamp::timestamp with time zone) * 1000) as timestamp_int FROM sensors_data WHERE sensor_id = $1 AND sensor_type = $2 ORDER BY id DESC LIMIT 1440;", sensorId, sensorType)
	if err != nil {
		fmt.Println("get Data: ", err)
		return nil, err
	}
	return data, nil
}


func GetSensors() ([]types.Sensor, error) {
	data := make([]types.Sensor, 0)
	err := DBX.Select(&data, "SELECT * FROM sensors;")
	if err != nil {
		fmt.Println("get Sensors: ", err)
		return nil, err
	}
	return data, nil
}
