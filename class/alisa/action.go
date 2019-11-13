package alisa

import (
	"../../db/postgres"
	"../../webserver/handlers"
	"../psh_devices"
	"../sonoff"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Action(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		fmt.Println("error: no body")
		handlers.HandlerError(w, "No body")
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Println("error: " + err.Error())
		handlers.HandlerError(w, err.Error())
		return
	}

	var t actionRequest

	err = json.Unmarshal(b, &t)
	if err != nil {
		fmt.Println()
		fmt.Println("error: " + err.Error())
		handlers.HandlerError(w, err.Error())
		return
	}

	fmt.Printf("%+v\n", t)

	//action

	for j := 0; j < len(t.Payload.Devices); j++ {
		device := t.Payload.Devices[j]

		dbDevice, err := postgres.GetDevice(device.ID)
		if err != nil {
			handlers.HandlerError(w, "1")
			return
		}

		byteValue, _ := json.Marshal(device.Capabilities[0].State.Value)

		switch dbDevice.Type {
		case "sonoff":

			for i := 0; i < len(device.Capabilities); i++ {

				switch device.Capabilities[i].Type {
				case "devices.capabilities.on_off":

					var val bool
					_ = json.Unmarshal(byteValue, &val)

					if val {
						sonoff.Switch("on", device.Capabilities[i].Type, device.ID)
					} else {
						sonoff.Switch("off", device.Capabilities[i].Type, device.ID)
					}
				}

			}

		case "psh-rgb":

			for i := 0; i < len(device.Capabilities); i++ {

				switch device.Capabilities[i].Type {
				case "devices.capabilities.on_off":

				case "devices.capabilities.color_setting":

					var val int64
					_ = json.Unmarshal(byteValue, &val)
					psh_devices.SetColor(device.ID, device.Capabilities[i].Type, val)

				}

			}

		}

	}

	// answer

	caps := make([]Capabilitie, 0)
	caps = append(caps, Capabilitie{
		Type: "devices.capabilities.on_off",
		State: State{
			Instance: t.Payload.Devices[0].Capabilities[0].State.Instance,
			ActionResult: ActionResult{
				Status: "DONE",
			},
		},
	})

	devices := make([]Device, 0)
	devices = append(devices, Device{
		ID:           t.Payload.Devices[0].ID,
		Capabilities: caps,
	})

	ans := CreateDeviceAnswer(r.Header.Get("X-Request-Id"), devices)
	handlers.HandlerInterfaceAssistant(w, ans)

}
