package types

type Sensor struct {
	SensorID        string       `json:"sensor_id"        db:"id"`
	Comment         string       `json:"comment"          db:"comment"`
	Place           string       `json:"place"            db:"place"`
	Alive           int64        `json:"alive"            db:"alive"`
	Data            []SensorData `json:"data"`
	RequestID       string       `json:"request_id"       db:"request_id"`
	Enable          bool         `json:"enable"           db:"enable"`
	UpdateTimestamp string       `json:"update_timestamp" db:"update_timestamp"`
}

type SensorData struct {
	ID              int64   `json:"id"                db:"id"`
	SensorID        string  `json:"sensor_id"         db:"sensor_id"`
	Type            int     `json:"type"              db:"sensor_type"`
	Value           float64 `json:"value"             db:"sensor_value"`
	TimeStampInt    float64 `json:"timestamp_int"     db:"timestamp_int"`
	TimeStampFormat string  `json:"timestamp_format"  db:"timestamp"`
	RequestID       string  `json:"request_id"        db:"request_id"`
}
