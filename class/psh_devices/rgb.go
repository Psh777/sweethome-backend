package psh_devices

import (
	"../../db/postgres"
	"../../modules/http_request"
	"fmt"
)
import "gopkg.in/go-playground/colors.v1"

func SetColor(deviceID string, setColor int64) {

	device, err := postgres.GetDevice(deviceID)

	if err != nil {
		fmt.Println(err)
		return
	}

	hexColor := fmt.Sprintf("%x", setColor)
	color, err := colors.Parse(hexColor)
	col := color.ToRGB()

	_, err = http_request.POST(device.Url, "/led?r="+fmt.Sprint(col.R)+"&g="+fmt.Sprint(col.G)+"&b"+fmt.Sprint(col.B), "")

	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
