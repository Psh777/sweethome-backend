package postgres

import (
	"../../types"
	"fmt"
)

func GetDevices() ([]types.DBDevice, error) {
	data := make([]types.DBDevice, 0)
	err := DBX.Select(&data, "SELECT * FROM devices;")
	if err != nil {
		fmt.Println("get Devices: ", err)
		return nil, err
	}
	return data, nil
}

func GetDevice(id string) (types.DBDevice, error) {
	var data types.DBDevice
	err := DBX.Get(&data, "SELECT * FROM devices WHERE id = $1;", id)
	if err != nil {
		fmt.Println("get Devices: ", err)
		return types.DBDevice{}, err
	}
	return data, nil
}
