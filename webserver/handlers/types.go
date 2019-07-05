package handlers


type Success struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
}

type Error struct {
	Success   bool        `json:"success"`
	ErrorCode string      `json:"code"`
	ErrorText string      `json:"text"`
	Result    interface{} `json:"result"`
}
