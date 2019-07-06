package types

type Sensor struct {
	SensorID  string       `json:"sensor_id"    db:"sensor_id"`
	Alive     int64        `json:"alive"        db:"alive"`
	Data      []SensorData `json:"data"`
	RequestID string       `json:"request_id"   db:"request_id"`
}

type SensorData struct {
	ID              int64   `json:"id"                db:"id"`
	SensorID        string  `json:"sensor_id"         db:"sensor_id"`
	Type            int     `json:"type"              db:"sensor_type"`
	Value           float64 `json:"value"             db:"sensor_value"`
	TimeStampInt    float64   `json:"timestamp_int"     db:"timestamp_int"`
	TimeStampFormat string  `json:"timestamp_format"  db:"timestamp"`

	RequestID string `json:"request_id"    db:"request_id"`
}
