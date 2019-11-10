package alisa

import (
	"../../webserver/handlers"
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

	// answer

	caps := make([]Capabilitie, 0)
	caps = append(caps, Capabilitie{
		Type: "devices.capabilities.on_off",
		State: State{
			Instance: "on",
			ActionResult: ActionResult{
				Status: "DONE",
			},
		},
	})

	devices := make([]Device, 0)
	devices = append(devices, Device{
		ID: "1",
	})

	payload := Payload{
		UserID:  "psh",
		Devices: devices,
	}

	answer := actionAnswer{
		RequestID: r.Header.Get("X-Request-Id"),
		Payload:   payload,
	}

	handlers.HandlerInterfaceAssistant(w, answer)

}

type actionRequest struct {
	Payload Payload `json:"payload"`
}

type actionAnswer struct {
	RequestID string  `json:"request_id"`
	Payload   Payload `json:"payload"`
}

type State struct {
	Instance     string       `json:"instance"`
	ActionResult ActionResult `json:"action_result"`
}

type ActionResult struct {
	Status string `json:"status"`
}
