// internal/config/config.go
package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURI  string
	DatabaseName string
}

var AppConfig *Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	AppConfig = &Config{
		DatabaseURI:  viper.GetString("DATABASE_URI"),
		DatabaseName: viper.GetString("DATABASE_NAME"),
	}
}
