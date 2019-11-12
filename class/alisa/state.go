package alisa

import (
	"../../db/postgres"
	"../../webserver/handlers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func DeviceState(w http.ResponseWriter, r *http.Request) {

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

	fmt.Printf("%+v\n", string(b))

	var t Payload

	err = json.Unmarshal(b, &t)
	if err != nil {
		fmt.Println()
		fmt.Println("error: " + err.Error())
		handlers.HandlerError(w, err.Error())
		return
	}

	fmt.Printf("%+v\n", t)

	//action
	
	devices := make([]Device, 0) 

	for j := 0; j < len(t.Devices); j++ {
		dbDevice, err := postgres.GetDevice(t.Devices[j].ID)
		if err != nil {
			handlers.HandlerError(w, err.Error())
			return
		}

		caps := make([]Capabilitie, 0)

		switch dbDevice.AlisaCapabilities {
		case "devices.capabilities.on_off":
			var v bool
			if dbDevice.State == "on" {
				v = true
			} else {
				v = false
			}
			state := State{
				Instance: "on",
				Value:    v,
			}
			caps = append(caps, Capabilitie{
				Type:       dbDevice.AlisaCapabilities,
				State:      state,
			})

		case "devices.capabilities.color_setting":
			//1
			var v bool
			if dbDevice.State == "on" {
				v = true
			} else {
				v = false
			}
			state := State{
				Instance: "on",
				Value:    v,
			}
			caps = append(caps, Capabilitie{
				Type:       dbDevice.AlisaCapabilities,
				State:      state,
			})
			//2
			state = State{
				Instance:     "rgb",
				Value:        0,
			}

			caps = append(caps, Capabilitie{
				Type:       dbDevice.AlisaCapabilities,
				State:      state,
			})
		}

		devices = append(devices, CreateDevice(dbDevice, caps))

	}

	ans := CreateDeviceAnswer(r.Header.Get("X-Request-Id"), devices)
	bb, _:=json.Marshal(ans)
	fmt.Println(string(bb))
	handlers.HandlerInterfaceAssistant(w, ans)

}
