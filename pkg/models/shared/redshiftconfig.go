package shared

type RedshiftConfig struct {
	DatabaseName string `json:"database_name"`
	Host         string `json:"host"`
	Password     string `json:"password"`
	Port         string `json:"port"`
	Username     string `json:"username"`
}
