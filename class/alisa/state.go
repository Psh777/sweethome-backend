package alisa

import (
	"../../db/postgres"
	"../../webserver/handlers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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

		dbCaps, err := postgres.GetCapabilities(t.Devices[j].ID)
		if err != nil {
			handlers.HandlerError(w, err.Error())
			return
		}

		caps := make([]Capabilitie, 0)

		for i := 0; i < len(dbCaps); i++ {
			switch dbCaps[i].Type {
			case "devices.capabilities.on_off":
				var v bool
				if dbCaps[i].State == "on" {
					v = true
				} else {
					v = false
				}
				state := State{
					Instance: dbCaps[i].Instance,
					Value:    v,
				}
				caps = append(caps, Capabilitie{
					Type:  "devices.capabilities.on_off",
					State: state,
				})

			case "devices.capabilities.color_setting":
				intState, err := strconv.ParseInt(dbCaps[i].State, 10, 64)
				if err != nil {
					fmt.Println(err)
					return
				}
				state := State{
					Instance: dbCaps[i].Instance,
					Value:    intState,
				}
				caps = append(caps, Capabilitie{
					Type:  "devices.capabilities.color_setting",
					State: state,
				})
			}
		}

		devices = append(devices, CreateDevice(dbDevice, caps))

	}

	ans := CreateDeviceAnswer(r.Header.Get("X-Request-Id"), devices)
	bb, _ := json.Marshal(ans)
	fmt.Println(string(bb))
	handlers.HandlerInterfaceAssistant(w, ans)

}
