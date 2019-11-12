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

		caps := make([]Capabilitie, 0)

		var par interface{}

		switch dbDevises[j].AlisaCapabilities {
		case "devices.capabilities.on_off":
		case "devices.capabilities.color_setting":
			par = Parameters{
				ColorModel:   "rgb",
				TemperatureK: TemperatureK{
					Min: 2700,
					Max: 9000,
				},
			}
		}

		caps = append(caps, Capabilitie{
			Type:       dbDevises[j].AlisaCapabilities,
			Parameters: par,
		})

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
	Type       string      `json:"type"`
	State      State       `json:"state"`
	Parameters interface{} `json:"parameters"`
}

type Parameters struct {
	ColorModel   string       `json:"color_model"`
	TemperatureK TemperatureK `json:"temperature_k"`
}

type TemperatureK struct {
	Min       int `json:"min"`
	Max       int `json:"max"`
	Precision int `json:"precision"`
}
