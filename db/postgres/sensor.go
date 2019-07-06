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

func GetData(sensorId string) ([]types.SensorData, error) {
	fmt.Println("get data: ", sensorId)
	data := make([]types.SensorData, 0)
	err := DBX.Select(&data, "SELECT *, ROUND(extract(epoch from timestamp::timestamp with time zone) * 1000) as timestamp_int FROM sensors_data WHERE sensor_id = $1 ORDER BY id DESC LIMIT 100;", sensorId)
	if err != nil {
		fmt.Println("get Data: ", err)
		return nil, err
	}
	return data, nil
}