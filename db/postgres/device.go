package postgres

import (
	"../../types"
	"fmt"
)

func GetDivices() ([]types.DBDevice, error) {
	data := make([]types.DBDevice, 0)
	err := DBX.Select(&data, "SELECT * FROM devices;")
	if err != nil {
		fmt.Println("get Devices: ", err)
		return nil, err
	}
	return data, nil
}

