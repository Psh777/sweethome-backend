package types

type DBDevice struct {
	ID                string `json:"id"          db:"id"`
	Type              string `json:"type"        db:"type"`
	AlisaType         string `json:"alisa_type"  db:"alisa_type"`
	Room              string `json:"room"        db:"room"`
	Name              string `json:"name"        db:"name"`
	Description       string `json:"description" db:"description"`
	Url               string `json:"url"         db:"url"`
}

type DBCapabilities struct {
	ID      int64  `json:"id"         db:"id"`
	Device  string `json:"device_id"  db:"device_id"`
	Type    string `json:"type"       db:"type"`
	Instans string `json:"instans"    db:"instans"`
	State   string `json:"state"      db:"state"`
}
