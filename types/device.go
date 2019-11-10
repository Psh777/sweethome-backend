package types

type DBDevice struct {
	ID                string `json:"id"          db:"id"`
	Type              string `json:"type"        db:"type"`
	IP                string `json:"ip"          db:"ip"`
	State             string `json:"state"       db:"state"`
	AlisaType         string `json:"alisa_type"  db:"alisa_type"`
	AlisaCapabilities string `json:"alisa_caps"  db:"alisa_caps"`
	Room              string `json:"room"        db:"room"`
	Name              string `json:"name"        db:"name"`
	Description       string `json:"description" db:"description"`
}
