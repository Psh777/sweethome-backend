package types

type MyConfig struct {
	Env        Env
	MainConfig MainConfig
}

type Env struct {
	PostgresHost        string `json:"postgres_host"`
	PostgresPort        string `json:"postgres_port"`
	PostgresUser        string `json:"postgres_user"`
	PostgresPassword    string `json:"postgres_password"`
	PostgresBase        string `json:"postgres_base"`
	HttpPort            string `json:"http_port"`
	TelegramBot         string `json:"telegram_bot"`
	SecurityBackend     string `json:"security_backend"`
	DialogFlowProjectID string `json:"dialog_flow_project_id"`
}

type MainConfig struct {
	ProjectName    string `json:"project_name"`
	ProjectUrl     string `json:"project_url"`
	PaginatorLimit int64  `json:"paginator_limit"`
	AdminPassword  string `json:"admin_password"`
}
