package config

import (
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var config = new(Config)

type Config struct {
	Mysql   MySQL   `toml:"mysql"`
	Session Session `toml:"session"`
	Redis   Redis   `toml:"redis"`
	Web     Web     `toml:"web"`
	Email   Email   `toml:"email"`
}

func init() {
	env := os.Getenv("TODO_ENV")
	if env == "" {
		env = "dev"
		log.Println("not set TODO_ENV the default mode is dev")
	} else {
		log.Println(env)
	}
	viper.SetConfigName(env + "_configs")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("配置已改变：%s", e.Name)
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})
}

func Get() *Config {
	return config
}
