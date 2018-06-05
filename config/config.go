package config

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
)

var (
	App AppConfig
)

func BootConfig() {
	App = AppConfig{}
	initProfile()
	initAppConfig()
}

func initProfile() {
	viper.SetDefault("profile", "dev")
	viper.BindEnv("profile")
	fmt.Println(viper.Get("profile"))
}

func initAppConfig() {
	configFileName := fmt.Sprintf("application-%s", viper.GetString("profile"))
	viper.SetConfigName(configFileName)
	viper.AddConfigPath(".")
	loadConfigFile()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		loadConfigFile()
	})
}

func loadConfigFile(){
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.UnmarshalKey("app", &App)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %+v", err))
	}
}

