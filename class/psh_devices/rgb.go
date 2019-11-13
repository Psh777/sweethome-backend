package psh_devices

import (
	"../../db/postgres"
	"../../modules/http_request"
	"fmt"
	"image/color"
)

func SetColor(deviceID string, setColor int64) {

	device, err := postgres.GetDevice(deviceID)

	if err != nil {
		fmt.Println(err)
		return
	}

	hex := fmt.Sprintf("#%06x", setColor)
	c, err := ParseHexColor(hex)
	if err != nil {
		return
	}
	a := "led?r="+fmt.Sprintf("%v",c.R)+"&g="+fmt.Sprint(c.G)+"&b="+fmt.Sprint(c.B)
	fmt.Print(a, err)

	_, err = http_request.GET(device.Url, a)

	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")

	}
	return
}