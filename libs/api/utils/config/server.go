package config

import (
	"os"
)

type ServerConfig struct {
	Port string
}

func GetServerConfig() ServerConfig {
	return ServerConfig{
		Port: os.Getenv("PORT"),
	}
}
