package bootstrap

import (
	"fmt"
	"gin-g/config"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func ParseConfig(path string) *viper.Viper {
	v := viper.New()
	v.SetConfigFile(path)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Read config file error: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		println(fmt.Sprintf("config file changed: $s", e.Name))
		println(fmt.Sprintf("config:%+v", config.Config()))
		if err := v.Unmarshal(config.Config()); err != nil {
			fmt.Println(fmt.Errorf("Parse config file on changed error: %s \n", err))
		}
	})
	if err := v.Unmarshal(config.Config()); err != nil {
		fmt.Println(fmt.Errorf("Parse config file error: %s \n", err))
	}
	println(fmt.Sprintf("config:%+v", config.Config()))
	
	return v
}
