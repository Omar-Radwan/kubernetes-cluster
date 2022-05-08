package config


type ApplicationConfig struct{
	Server  ServerConfig  `json:"server,omitempty"`
	Database DatabaseConfig `json:"database,omitempty"`
}

type ServerConfig struct {
	Host         string `json:"host,omitempty"`
	Port         int    `json:"port,omitempty"`
	ReadTimeout  int    `json:"readTimeout,omitempty"`
	WriteTimeout int    `json:"writeTimeout"`
}


type DatabaseConfig struct {
	DatabaseName string `json:"databaseName,omitempty"`
	Host   string `json:"host,omitempty"`
	Port   int    `json:"port,omitempty"`
}
