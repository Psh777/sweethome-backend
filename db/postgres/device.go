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
	fmt.Println("Get device:", id)
	var data types.DBDevice
	err := DBX.Get(&data, "SELECT * FROM devices WHERE id = $1;", id)
	if err != nil {
		fmt.Println("get Devices: ", err)
		return types.DBDevice{}, err
	}
	return data, nil
}

func SetState(id, state string) error {
	fmt.Println("Set State device:", id, state)
	_, err := DBX.Exec("UPDATE device SET state = $1 WHERE id = $2;", state, id)
	if err != nil {
		fmt.Println("postgres UpdateState: ", err)
		return err
	}
	return nil
}
