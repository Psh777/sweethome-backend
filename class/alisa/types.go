package alisa

type actionRequest struct {
	Payload Payload `json:"payload"`
}

type actionAnswer struct {
	RequestID string  `json:"request_id"`
	Payload   Payload `json:"payload"`
}

type State struct {
	Instance     string       `json:"instance,omitempty"`
	Value        interface{}  `json:"value,omitempty"`
	ActionResult ActionResult `json:"action_result,omitempty"`
}

type ActionResult struct {
	Status string `json:"status,omitempty"`
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
	Name         string        `json:"name,omitempty"`
	Description  string        `json:"description,omitempty"`
	Room         string        `json:"room,omitempty"`
	Type         string        `json:"type,omitempty"`
	Capabilities []Capabilitie `json:"capabilities"`
}

type Capabilitie struct {
	Type       string      `json:"type"`
	State      State       `json:"state,omitempty"`
	Parameters interface{} `json:"parameters,omitempty"`
}

type Parameters struct {
	ColorModel string `json:"color_model,omitempty"`
	//TemperatureK TemperatureK `json:"temperature_k"`
	Value int64 `json:"value,omitempty"`
}

type TemperatureK struct {
	Min       int `json:"min"`
	Max       int `json:"max"`
	Precision int `json:"precision"`
}
