package config

import (
	"os"
)

const (
	Port         = "PORT"
	Username     = "USERNAME"
	Password     = "PASSWORD"
	DatabaseName = "DATABASE_NAME"
	Host         = "HOST"
)

type Config struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DatabaseName string
}

func Get() Config {
	return Config{
		Host:         os.Getenv(Host),
		Port:         os.Getenv(Port),
		Username:     os.Getenv(Username),
		Password:     os.Getenv(Password),
		DatabaseName: os.Getenv(DatabaseName),
	}
}
