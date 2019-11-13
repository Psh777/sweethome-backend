package alisa

import (
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

	byteValue, _ := json.Marshal(t.Payload.Devices[0].Capabilities[0].State.Value)

	switch t.Payload.Devices[0].Capabilities[0].Type {
	case "devices.capabilities.on_off":

		var val bool
		_ = json.Unmarshal(byteValue, &val)

		if val {
			sonoff.Switch("on", t.Payload.Devices[0].ID)
		} else {
			sonoff.Switch("off", t.Payload.Devices[0].ID)
		}



	case "devices.capabilities.color_setting":

		var val int64
		_ = json.Unmarshal(byteValue, &val)
		psh_devices.SetColor(t.Payload.Devices[0].ID, val)

	default:
		return
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
