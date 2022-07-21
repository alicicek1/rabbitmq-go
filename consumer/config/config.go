package config

import (
	"log"
	"os"
)

type Configuration struct {
	RabbitUsername string
	RabbitPassword string
	RabbitUri      string
	ErrorQName     string
}

func GetConfigs() map[string]Configuration {

	return AppConfiguration{
		"qa": Configuration{
			RabbitUsername: "guest",
			RabbitPassword: "guest",
			RabbitUri:      "localhost:5672",
			ErrorQName:     "error",
		},
		"prod": Configuration{
			RabbitUsername: "prod",
			RabbitPassword: "prod",
			RabbitUri:      "localhost:5672",
			ErrorQName:     "error",
		},
	}
}

type AppConfiguration map[string]Configuration

func GetConfig() Configuration {
	env := os.Getenv("Env")
	if env == "" {
		log.Panicf("%s", "Environment variable is not found.")
	}
	cfg, ok := GetConfigs()[env]
	if !ok {
		log.Panicf("%s", "Config data is not found.")
	}
	return cfg
}
