package alisa

import (
	"../../webserver/handlers"
	"net/http"
)

func Devices(w http.ResponseWriter, r *http.Request) {

	caps := make([]Capabilitie, 0)
	caps = append(caps, Capabilitie{Type: "devices.capabilities.on_off"})

	device := Device{
		ID:           "1",
		Name:         "лампa",
		Description:  "лампa",
		Room:         "кабинет",
		Type:         "devices.types.light",
		Capabilities: caps,
	}

	devices := make([]Device, 0)
	devices = append(devices, device)

	payload := Payload{
		UserID:  "psh",
		Devices: devices,
	}

	answer := Answer{
		RequestID: r.Header.Get("X-Request-Id"),
		Payload:   payload,
	}

	handlers.HandlerInterfaceAssistant(w, answer)

}

type Answer struct {
	RequestID string  `json:"request_id"`
	Payload   Payload `json:"payload"`
}

type Payload struct {
	UserID  string   `json:"user_id"`
	Devices []Device `json:"devices"`
}

type Device struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Room         string        `json:"room"`
	Type         string        `json:"type"`
	Capabilities []Capabilitie `json:"capabilities"`
}

type Capabilitie struct {
	Type  string `json:"type"`
	State State  `json:"state"`
}
