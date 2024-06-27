package config

import (
    "github.com/spf13/viper"
	"log"
)

func LoadConfig(){
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration file: %v", err)
	}

}
