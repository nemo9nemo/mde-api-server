package common

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int
	}

	Database struct {
		Host     string
		Port     int
		Username string
		Password string
		Name     string
	}

	Redis struct {
		Host string
		Port int
	}
}

var Cfg Config

func LoadConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("[Error] No config file found: %v", err)
	}
	if err := viper.Unmarshal(&Cfg); err != nil {
		log.Fatalf("[Error] Config unmarshal failed: %v", err)
	}

	fmt.Printf("[OK] Config loaded. Server Port: %d\n", Cfg.Server.Port)
}
