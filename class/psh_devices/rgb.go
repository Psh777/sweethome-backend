package psh_devices

import (
	"../../db/postgres"
	"../../modules/http_request"
	"fmt"
)

func SetColor(deviceID string, h, s, v int) {

	device, err := postgres.GetDevice(deviceID)

	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = http_request.POST(device.Url, "/led?h="+fmt.Sprint(h)+"&s="+fmt.Sprint(h)+"&v"+fmt.Sprint(v), "")

	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
