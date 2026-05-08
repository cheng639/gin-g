package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var serviceConfig ServiceConfig

func Config() *ServiceConfig {
	return &serviceConfig
}

type ServiceConfig struct {
	WorkDir      string
	Server       Server       `json:"server" yaml:"server"`
	Logger       LoggerConfig `json:"logger" yaml:"logger"`
	LoggerWriter LoggerWriter `json:"loggerWriter" yaml:"loggerwriter"`
	Redis        RedisConfig  `json:"redis" yaml:"redis"`
}

type Server struct {
	Name string `json:"name" yaml:"name"`
	Host string `json:"host" yaml:"host"`
	IP   string `json:"ip" yaml:"ip"`
	Port string `json:"port" yaml:"port"`
}

func ParseConfig(path string) *viper.Viper {
	v := viper.New()
	v.SetConfigFile(path)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Read config file error: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		println(fmt.Sprintf("config file changed: %s", e.Name))
		println(fmt.Sprintf("config:%+v", Config()))
		if err2 := v.Unmarshal(Config()); err2 != nil {
			fmt.Println(fmt.Errorf("Parse config file on changed error: %s \n", err2))
		}
	})
	if err1 := v.Unmarshal(Config()); err1 != nil {
		fmt.Println(fmt.Errorf("Parse config file error: %s \n", err1))
	}
	println(fmt.Sprintf("config:%+v", Config()))

	return v
}
