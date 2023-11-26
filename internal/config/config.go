package config

import (
	"github.com/spf13/viper"
	"log"
)

func InitConfig() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file: %s", err)
	}
}
