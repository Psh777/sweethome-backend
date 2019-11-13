package psh_devices

import (
	"../../db/postgres"
	"../../modules/http_request"
	"fmt"
)

import "github.com/lucasb-eyer/go-colorful"

func SetColor(deviceID string, setColor int64) {

	device, err := postgres.GetDevice(deviceID)

	if err != nil {
		fmt.Println(err)
		return
	}

	hexColor := fmt.Sprintf("#%06x", setColor)
	c, err := colorful.Hex(hexColor)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = http_request.POST(device.Url, "led?r="+fmt.Sprint(c.R)+"&g="+fmt.Sprint(c.G)+"&b"+fmt.Sprint(c.B), "")

	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
