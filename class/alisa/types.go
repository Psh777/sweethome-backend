package alisa

type actionRequest struct {
	Payload Payload `json:"payload"`
}

type actionAnswer struct {
	RequestID string  `json:"request_id"`
	Payload   Payload `json:"payload"`
}

type State struct {
	Instance     string       `json:"instance"`
	Value        interface{}  `json:"value"`
	ActionResult ActionResult `json:"action_result"`
}

type ActionResult struct {
	Status string `json:"status"`
}

type HSV struct {
	H int `json:"h"`
	S int `json:"s"`
	V int `json:"v"`
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
	Value        int64        `json:"value"`
}

type TemperatureK struct {
	Min       int `json:"min"`
	Max       int `json:"max"`
	Precision int `json:"precision"`
}