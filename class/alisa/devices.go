package alisa

import (
	"../../db/postgres"
	"../../webserver/handlers"
	"fmt"
	"net/http"
	"strconv"
)

func Devices(w http.ResponseWriter, r *http.Request) {

	dbDevises, err := postgres.GetDevices()

	if err != nil {
		fmt.Println(err)
		return
	}

	devices := make([]Device, 0)

	for j := 0; j < len(dbDevises); j++ {

		dbCaps, err := postgres.GetCapabilities(dbDevises[j].ID)
		if err != nil {
			fmt.Println(err)
			return
		}

		var par interface{}
		caps := make([]Capabilitie, 0)

		for i := 0; i < len(dbCaps); i++ {

			switch dbCaps[i].Type {
			case "devices.capabilities.on_off":

				caps = append(caps, Capabilitie{
					Type: "devices.capabilities.on_off",
				})

			case "devices.capabilities.color_setting":

				intState, err := strconv.ParseInt(dbCaps[i].State, 10, 64)
				if err != nil {
					fmt.Println(err)
					return
				}
				par = Parameters{
					ColorModel: dbCaps[i].Instans,
					Value: intState,
				}
				caps = append(caps, Capabilitie{
					Type:       "devices.capabilities.color_setting",
					Parameters: par,
				})

			}

		}

		device := CreateDevice(dbDevises[j], caps)
		devices = append(devices, device)

	}

	ans := CreateDeviceAnswer(r.Header.Get("X-Request-Id"), devices)
	handlers.HandlerInterfaceAssistant(w, ans)

}
