package types

type Sensor struct {
	SensorID  string       `json:"sensor_id"    db:"sensor_id"`
	Alive     int64        `json:"alive"        db:"alive"`
	Data      []SensorData `json:"data"`
	RequestID string       `json:"request_id"    db:"request_id"`
}

type SensorData struct {
	ID        int64       `json:"id"            db:"id"`
	SensorID  string      `json:"sensor_id"     db:"sensor_id"`
	Type      int         `json:"type"          db:"sensor_type"`
	Value     interface{} `json:"value"         db:"value"`
	RequestID string      `json:"request_id"    db:"request_id"`
}
