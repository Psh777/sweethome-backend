package psh_devices

import (
	"fmt"
	"github.com/Psh777/sweethome-backend/db/postgres"
	"github.com/Psh777/sweethome-backend/modules/http_request"
	"image/color"
	"strconv"
)

func Switch(deviceID string, capabilities string, state string) {

	device, err := postgres.GetDevice(deviceID)
	if err != nil {
		fmt.Println(err)
		return
	}

	if state == "off" {

		c, err := ParseHexColor("#000000")
		if err != nil {
			fmt.Println(err)
			return
		}

		a := "led?r=" + fmt.Sprintf("%v", c.R) + "&g=" + fmt.Sprint(c.G) + "&b=" + fmt.Sprint(c.B)
		fmt.Println(a, err)

		_, err = http_request.GET(device.Url, a)
		if err != nil {
			fmt.Println(err)
			return
		}

	} else {

		caps, _ := postgres.GetCapabilities(deviceID)
		for i := 0; i < len(caps); i++ {
			if caps[i].Instance == "rgb" {
				intState, err := strconv.ParseInt(caps[i].State, 10, 64)
				if err != nil {
					fmt.Println(err)
					return
				}
				SetColor(deviceID, capabilities, intState)
			}
		}

	}

	_ = postgres.SetState(deviceID, capabilities, state)

}

func SetColor(deviceID string, capabilities string, setColor int64) {

	device, err := postgres.GetDevice(deviceID)

	if err != nil {
		fmt.Println(err)
		return
	}

	hex := fmt.Sprintf("#%06x", setColor)
	c, err := ParseHexColor(hex)
	if err != nil {
		fmt.Println(err)
		return
	}
	a := "color?r=" + fmt.Sprintf("%v", c.R) + "&g=" + fmt.Sprint(c.G) + "&b=" + fmt.Sprint(c.B)
	fmt.Println(a, err)

	_, err = http_request.GET(device.Url, a)
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = postgres.SetState(deviceID, capabilities, fmt.Sprint(setColor))

	return
}

func SetBrightness(deviceID string, capabilities string, data int64) {

	device, err := postgres.GetDevice(deviceID)

	if err != nil {
		fmt.Println(err)
		return
	}

	a := "brightness?range=" + fmt.Sprintf("%v", data)
	fmt.Println(a, err)

	_, err = http_request.GET(device.Url, a)
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = postgres.SetState(deviceID, capabilities, fmt.Sprint(data))

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
