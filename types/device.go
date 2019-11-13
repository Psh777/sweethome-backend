package types

type DBDevice struct {
	ID                string `json:"id"          db:"id"`
	Type              string `json:"type"        db:"type"`
	IP                string `json:"ip"          db:"ip"`
	State1            string `json:"state1"      db:"state1"`
	State2            string `json:"state2"      db:"state2"`
	State3            string `json:"state3"      db:"state3"`
	AlisaType         string `json:"alisa_type"  db:"alisa_type"`
	AlisaCapabilities string `json:"alisa_caps"  db:"alisa_caps"`
	Room              string `json:"room"        db:"room"`
	Name              string `json:"name"        db:"name"`
	Description       string `json:"description" db:"description"`
	Url               string `json:"url"         db:"url"`
}
