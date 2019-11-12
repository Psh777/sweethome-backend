package alisa

import (
	"../../db/postgres"
	"../../webserver/handlers"
	"fmt"
	"net/http"
)

func Devices(w http.ResponseWriter, r *http.Request) {

	dbDevises, err := postgres.GetDevices()

	if err != nil {
		fmt.Println(err)
		return
	}

	devices := make([]Device, 0)

	for j := 0; j < len(dbDevises); j++ {

		var par interface{}
		caps := make([]Capabilitie, 0)

		switch dbDevises[j].AlisaCapabilities {
		case "devices.capabilities.on_off":

			caps = append(caps, Capabilitie{
				Type:       dbDevises[j].AlisaCapabilities,
				Parameters: par,
			})
		case "devices.capabilities.color_setting":
			//1
			caps = append(caps, Capabilitie{
				Type:       dbDevises[j].AlisaCapabilities,
				Parameters: par,
			})
			//2
			par = Parameters{
				ColorModel: "rgb",
				Value:      0,
			}
			caps = append(caps, Capabilitie{
				Type:       dbDevises[j].AlisaCapabilities,
				Parameters: par,
			})
		}

		device := CreateDevice(dbDevises[j], caps)
		devices = append(devices, device)

	}

	ans := CreateDeviceAnswer(r.Header.Get("X-Request-Id"), devices)
	handlers.HandlerInterfaceAssistant(w, ans)

}
