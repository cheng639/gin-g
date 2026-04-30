package config

var serviceConfig ServiceConfig

func Config() *ServiceConfig {
	return &serviceConfig
}

type ServiceConfig struct {
	WorkDir      string
	Server       Server       `json:"server" yaml:"server"`
	Logger       LoggerConfig `json:"logger" yaml:"logger"`
	LoggerWriter LoggerWriter `json:"loggerWriter" yaml:"loggerwriter"`
	RedisConfig  RedisConfig  `json:"redis" yaml:"redis"`
}

type Server struct {
	Name string `json:"name" yaml:"name"`
	Host string `json:"host" yaml:"host"`
	IP   string `json:"ip" yaml:"ip"`
	Port string `json:"port" yaml:"port"`
}
