package alisa

import (
	"../../db/postgres"
	"../../webserver/handlers"
	"fmt"
	"net/http"
)

func Devices(w http.ResponseWriter, r *http.Request) {

	dbDevises, err := postgres.GetDivices()

	if err != nil {
		fmt.Println(err)
		return
	}

	devices := make([]Device, 0)

	for j := 0; j <= len(dbDevises); j++ {

		caps := make([]Capabilitie, 0)
		caps = append(caps, Capabilitie{Type: dbDevises[j].AlisaCapabilities})

		device := Device{
			ID:           dbDevises[j].ID,
			Name:         dbDevises[j].Name,
			Description:  dbDevises[j].Description,
			Room:         dbDevises[j].Room,
			Type:         dbDevises[j].AlisaType,
			Capabilities: caps,
		}

		devices = append(devices, device)

	}

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
